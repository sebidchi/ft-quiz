package domain

import "github.com/sebidchi/ft-quiz/internal/pkg/domain"

const errorStoring = "error storing answers"

type ErrorStoringAnswers struct {
	domain.CriticalError
	extraItems map[string]interface{}
}

func NewErrorStoringAnswers(userId string) *ErrorStoringAnswers {
	return &ErrorStoringAnswers{extraItems: map[string]interface{}{"userId": userId}}
}

func (q *ErrorStoringAnswers) ExtraItems() map[string]interface{} {
	return q.extraItems
}

func (q *ErrorStoringAnswers) Error() string {
	return errorStoring
}
