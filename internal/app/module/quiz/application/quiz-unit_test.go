package application_test

import (
	"context"

	"github.com/sebidchi/ft-quiz/internal/app/module/quiz/domain"
	"github.com/stretchr/testify/mock"
)

type QuizRepositoryMock struct {
	mock.Mock
}

func (q *QuizRepositoryMock) GetQuiz(ctx context.Context, id string) (*domain.Quiz, error) {
	args := q.Called(ctx, id)
	if args[0] == nil {
		return nil, args.Error(1)
	}
	return args[0].(*domain.Quiz), nil
}

func (q *QuizRepositoryMock) ShouldGetQuiz(ctx context.Context, id string, quiz *domain.Quiz, err error) {
	q.On("GetQuiz", ctx, id).Once().Return(quiz, err)
}
func (q *QuizRepositoryMock) ShouldGetQuizAndFail(ctx context.Context, id string, err error) {
	q.On("GetQuiz", ctx, id).Once().Return(nil, err)
}
