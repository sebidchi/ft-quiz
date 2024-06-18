package domain

import "context"

type UserAnswerRepository interface {
	SaveUserAnswers(ctx context.Context, userAnswers *UserAnswers) error
	GetUsersResults(ctx context.Context) (UsersResults, error)
}
