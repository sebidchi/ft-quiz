package test

import (
	"context"

	"github.com/sebidchi/ft-quiz/cmd/di"
	"github.com/sebidchi/ft-quiz/internal/app/module/user-answer/infrastructure"
)

type InMemoryDbArranger struct {
	di *di.FTQuizDI
}

func NewInMemoryDbArranger(di *di.FTQuizDI) *InMemoryDbArranger {
	return &InMemoryDbArranger{di: di}
}

func (db *InMemoryDbArranger) Arrange(_ context.Context) {
	db.di.UserAnswersServices.UserAnswerRepository = infrastructure.NewInMemoryUserAnswerRepository()
}
