package application_test

import (
	"context"
	"testing"

	"github.com/sebidchi/ft-quiz/internal/app/module/user-answer/application"
	"github.com/sebidchi/ft-quiz/internal/app/module/user-answer/domain"
	"github.com/sebidchi/ft-quiz/internal/pkg/domain/bus"
	"github.com/sebidchi/ft-quiz/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestStoreUserAnswersCommandHandler(t *testing.T) {
	repo := new(UserAnswersRepositoryMock)

	storer := domain.NewUserAnswersStorer(repo)
	handler := application.NewStoreUserAnswersCommandHandler(storer)

	userId := "test-user-1"
	answers := map[string]string{"question1": "option1", "question2": "option2"}
	userAnswers := &domain.UserAnswers{
		UserId:  userId,
		Answers: domain.NewUserAnswersFromRaw(answers),
	}

	tests := []struct {
		name         string
		command      bus.Command
		expectations func(ctx context.Context, repo *UserAnswersRepositoryMock, err error)
		wantedErr    error
	}{
		{
			name: "Alright !",
			command: application.NewStoreUserAnswersCommand(
				"test-user-1",
				answers,
			),
			expectations: func(ctx context.Context, repo *UserAnswersRepositoryMock, err error) {
				repo.ShouldSaveUserAnswers(ctx, userAnswers, err)
			},
			wantedErr: nil,
		},
		{
			name: "Error storing data :( ",
			command: application.NewStoreUserAnswersCommand(
				"test-user-1",
				answers,
			),
			expectations: func(ctx context.Context, repo *UserAnswersRepositoryMock, err error) {
				repo.ShouldSaveUserAnswers(ctx, userAnswers, err)
			},
			wantedErr: domain.NewErrorStoringAnswers(userId),
		},
		{
			name:    "Invalid Command",
			command: test.InvalidCommand{},
			expectations: func(ctx context.Context, repo *UserAnswersRepositoryMock, err error) {
			},
			wantedErr: bus.NewCommandNotValid(test.InvalidCommand{}.Id()),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.expectations(context.Background(), repo, tt.wantedErr)
			err := handler.Handle(context.Background(), tt.command)
			mock.AssertExpectationsForObjects(t, repo)

			if tt.wantedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.wantedErr, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
