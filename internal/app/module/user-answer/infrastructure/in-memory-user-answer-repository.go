package infrastructure

import (
	"context"
	"sync"

	"github.com/sebidchi/ft-quiz/internal/app/module/user-answer/domain"
)

type InMemoryUserAnswerRepository struct {
	userAnswers map[string]domain.Result
	mutex       sync.RWMutex
}

func NewInMemoryUserAnswerRepository() *InMemoryUserAnswerRepository {
	return &InMemoryUserAnswerRepository{userAnswers: make(map[string]domain.Result), mutex: sync.RWMutex{}}
}

func (i *InMemoryUserAnswerRepository) SaveUserAnswers(_ context.Context, userAnswers *domain.UserAnswers) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	i.userAnswers[userAnswers.UserId] = domain.Result{
		Points:         userAnswers.Score(),
		TotalQuestions: len(userAnswers.Answers),
	}

	return nil
}

func (i *InMemoryUserAnswerRepository) GetUsersResults(_ context.Context) (domain.UsersResults, error) {
	i.mutex.RLock()
	defer i.mutex.RUnlock()
	results := i.userAnswers
	return results, nil
}
