package application_test

import (
	"context"

	"github.com/sebidchi/ft-quiz/internal/app/module/user-answer/domain"
	"github.com/stretchr/testify/mock"
)

type UserAnswersRepositoryMock struct {
	mock.Mock
}

func (u *UserAnswersRepositoryMock) SaveUserAnswers(ctx context.Context, userAnswers *domain.UserAnswers) error {
	args := u.Called(ctx, userAnswers)
	return args.Error(0)
}

func (u *UserAnswersRepositoryMock) GetUsersResults(ctx context.Context) (domain.UsersResults, error) {
	args := u.Called(ctx)
	if args[0] == nil {
		return nil, args.Error(1)
	}
	return args[0].(domain.UsersResults), nil
}

func (u *UserAnswersRepositoryMock) ShouldSaveUserAnswers(ctx context.Context, userAnswers *domain.UserAnswers, err error) {
	u.On("SaveUserAnswers", ctx, userAnswers).Once().Return(err)
}

func (u *UserAnswersRepositoryMock) ShouldGetUsersResults(ctx context.Context, userResults domain.UsersResults, err error) {
	u.On("GetUsersResults", ctx).Once().Return(userResults, err)
}
