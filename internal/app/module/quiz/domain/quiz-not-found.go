package domain

import "github.com/sebidchi/ft-quiz/internal/pkg/domain"

const quizNotFound = "quiz not found"

type QuizNotFound struct {
	domain.DomainError
	extraItems map[string]interface{}
}

func NewQuizNotFound(id string) *QuizNotFound {
	return &QuizNotFound{extraItems: map[string]interface{}{"id": id}}
}

func (q *QuizNotFound) ExtraItems() map[string]interface{} {
	return q.extraItems
}

func (q *QuizNotFound) Error() string {
	return quizNotFound
}
