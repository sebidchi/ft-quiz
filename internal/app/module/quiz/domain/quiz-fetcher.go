package domain

import "context"

type QuizFetcher struct {
	quizRepository QuizRepository
}

func NewQuizFetcher(quizRepository QuizRepository) *QuizFetcher {
	return &QuizFetcher{quizRepository: quizRepository}
}

func (qf QuizFetcher) FetchQuiz(ctx context.Context, quizId string) (*Quiz, error) {
	quiz, err := qf.quizRepository.GetQuiz(ctx, quizId)
	if err != nil {
		return nil, err
	}

	return quiz, nil
}
