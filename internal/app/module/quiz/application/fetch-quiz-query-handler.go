package application

import (
	"context"

	"github.com/sebidchi/ft-quiz/internal/app/module/quiz/domain"
	"github.com/sebidchi/ft-quiz/internal/pkg/domain/bus"
)

type FetchQuizQuery struct {
	quizId string
}

func NewFetchQuizQuery(quizId string) *FetchQuizQuery {
	return &FetchQuizQuery{quizId: quizId}
}

func (f FetchQuizQuery) Id() string {
	return "FetchQuizQuery"
}

type FetchQuizQueryHandler struct {
	quizFetcher *domain.QuizFetcher
}

func NewFetchQuizQueryHandler(quizFetcher *domain.QuizFetcher) *FetchQuizQueryHandler {
	return &FetchQuizQueryHandler{quizFetcher: quizFetcher}
}

func (f FetchQuizQueryHandler) Handle(ctx context.Context, query bus.Query) (interface{}, error) {
	fetQuizQuery, ok := query.(*FetchQuizQuery)
	if !ok {
		return nil, bus.NewInvalidQuery(query)
	}

	res, err := f.quizFetcher.FetchQuiz(ctx, fetQuizQuery.quizId)
	if err != nil {
		return nil, err
	}

	return NewQuizResponseFromQuiz(res), nil
}
