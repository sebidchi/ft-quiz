package domain

import "context"

type UserResultsFetcher struct {
	repository UserAnswerRepository
}

func NewUserResultsFetcher(repository UserAnswerRepository) *UserResultsFetcher {
	return &UserResultsFetcher{repository: repository}
}

func (f *UserResultsFetcher) Fetch(ctx context.Context, userId string) (*UserResults, error) {
	res, err := f.repository.GetUsersResults(ctx)
	if err != nil {
		return nil, err
	}

	return res.UserResults(userId)
}
