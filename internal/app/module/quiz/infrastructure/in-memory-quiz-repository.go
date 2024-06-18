package infrastructure

import (
	"context"

	"github.com/sebidchi/ft-quiz/internal/app/module/quiz/domain"
)

// InMemoryQuizRepository struct is a repository that stores the quiz in memory
// and returns it when requested
type InMemoryQuizRepository struct {
	quiz map[string]*domain.Quiz
}

// NewInMemoryQuizRepository creates a new InMemoryQuizRepository
// with the default quiz
// We are giving a default quiz to the repository
// so that we can test the application without a database
func NewInMemoryQuizRepository(defaultQuiz *domain.Quiz) *InMemoryQuizRepository {
	quizzes := make(map[string]*domain.Quiz)
	quizzes["default"] = defaultQuiz
	return &InMemoryQuizRepository{
		quiz: quizzes,
	}
}

// GetQuiz returns the quiz with the given id from the repository
// if the quiz is not found, it returns a QuizNotFound error
func (r *InMemoryQuizRepository) GetQuiz(ctx context.Context, quizId string) (*domain.Quiz, error) {
	quiz, ok := r.quiz[quizId]
	if ok {
		return quiz, nil
	}

	return nil, domain.NewQuizNotFound(quizId)
}
