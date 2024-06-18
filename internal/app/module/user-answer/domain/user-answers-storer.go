package domain

import "context"

type UserAnswersStorer struct {
	UserAnswersRepository UserAnswerRepository
}

func NewUserAnswersStorer(userAnswersRepository UserAnswerRepository) *UserAnswersStorer {
	return &UserAnswersStorer{UserAnswersRepository: userAnswersRepository}
}

func (s UserAnswersStorer) Store(ctx context.Context, userId string, answers map[string]string) error {
	userAnswers := &UserAnswers{UserId: userId, Answers: NewUserAnswersFromRaw(answers)}
	return s.UserAnswersRepository.SaveUserAnswers(ctx, userAnswers)
}
