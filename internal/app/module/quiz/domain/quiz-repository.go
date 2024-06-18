package domain

import "context"

type QuizRepository interface {
	GetQuiz(ctx context.Context, id string) (*Quiz, error)
}
