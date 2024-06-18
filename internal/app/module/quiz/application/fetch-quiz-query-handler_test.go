package application_test

import (
	"context"
	"testing"

	"github.com/sebidchi/ft-quiz/internal/app/module/quiz/application"
	"github.com/sebidchi/ft-quiz/internal/app/module/quiz/domain"
	"github.com/sebidchi/ft-quiz/internal/pkg/domain/bus"
	"github.com/sebidchi/ft-quiz/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetchQuizQueryHandler(t *testing.T) {
	quiz1 := test.RandomQuizWithQuestions()
	repo := new(QuizRepositoryMock)

	handler := application.NewFetchQuizQueryHandler(domain.NewQuizFetcher(repo))

	tests := []struct {
		name           string
		query          bus.Query
		id             string
		wantedResponse *application.QuizResponse
		expectations   func(ctx context.Context, quizId string, quiz *domain.Quiz, repo *QuizRepositoryMock, err error)
		wantedError    error
	}{
		{
			name:           "Quiz 1",
			query:          application.NewFetchQuizQuery("default"),
			id:             "default",
			wantedResponse: application.NewQuizResponseFromQuiz(quiz1),
			expectations: func(ctx context.Context, quizId string, quiz *domain.Quiz, repo *QuizRepositoryMock, err error) {
				repo.ShouldGetQuiz(ctx, quizId, quiz, err)
			},
			wantedError: nil,
		},
		{
			name:           "Non-existent Quiz",
			query:          application.NewFetchQuizQuery("non-existent-id"),
			id:             "non-existent-id",
			wantedResponse: nil,
			expectations: func(ctx context.Context, quizId string, quiz *domain.Quiz, repo *QuizRepositoryMock, err error) {
				repo.ShouldGetQuizAndFail(ctx, quizId, err)
			},
			wantedError: domain.NewQuizNotFound("non-existent-id"),
		},
		{
			name:           "Invalid query",
			query:          test.InvalidQuery{},
			wantedResponse: nil,
			expectations: func(ctx context.Context, quizId string, quiz *domain.Quiz, repo *QuizRepositoryMock, err error) {
			},
			wantedError: bus.NewInvalidQuery(test.InvalidQuery{}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.expectations(context.Background(), tt.id, quiz1, repo, tt.wantedError)
			result, err := handler.Handle(context.Background(), tt.query)
			mock.AssertExpectationsForObjects(t, repo)

			if tt.wantedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.wantedError, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantedResponse, result)
			}
		})
	}
}
