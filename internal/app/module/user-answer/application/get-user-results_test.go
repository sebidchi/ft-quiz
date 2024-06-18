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

func TestGetUserResultsQueryHandler(t *testing.T) {
	repo := new(UserAnswersRepositoryMock)

	handler := application.NewGetUserResultsQueryHandler(domain.NewUserResultsFetcher(repo))

	userId := "aUserId"
	noExistingUserId := "non-existent-id"
	usersResults := domain.UsersResults{
		userId: domain.Result{
			Points:         5,
			TotalQuestions: 5,
		},
		"anotherUserId": domain.Result{
			Points:         5,
			TotalQuestions: 5,
		},
		"yetAnotherUserId": domain.Result{
			Points:         4,
			TotalQuestions: 5,
		},
	}
	userResults := &domain.UserResults{
		UserId:     userId,
		Total:      5,
		Percentage: 100,
		BetterThan: 50,
		TotalUsers: 3,
	}

	tests := []struct {
		name           string
		query          bus.Query
		id             string
		wantedResponse *application.UserResultsResponse
		expectations   func(ctx context.Context, userId string, res domain.UsersResults, repo *UserAnswersRepositoryMock, err error)
		wantedError    error
	}{
		{
			name:           "alright !",
			query:          application.NewGetUserResultsQuery(userId),
			id:             userId,
			wantedResponse: application.NewUserResultsResponse(userResults),
			expectations: func(ctx context.Context, userId string, res domain.UsersResults, repo *UserAnswersRepositoryMock, err error) {
				repo.ShouldGetUsersResults(ctx, res, err)
			},
			wantedError: nil,
		},
		{
			name:           "Non-existent User",
			query:          application.NewGetUserResultsQuery(noExistingUserId),
			id:             noExistingUserId,
			wantedResponse: nil,
			expectations: func(ctx context.Context, userId string, res domain.UsersResults, repo *UserAnswersRepositoryMock, err error) {
				repo.ShouldGetUsersResults(ctx, res, err)
			},
			wantedError: domain.NewUsersResultsNotFound(noExistingUserId),
		},
		{
			name:           "Invalid query",
			query:          test.InvalidQuery{},
			wantedResponse: nil,
			expectations: func(ctx context.Context, userId string, res domain.UsersResults, repo *UserAnswersRepositoryMock, err error) {
			},
			wantedError: bus.NewInvalidQuery(test.InvalidQuery{}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.expectations(context.Background(), tt.id, usersResults, repo, tt.wantedError)
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
