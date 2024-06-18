package test

import (
	"github.com/brianvoe/gofakeit"
	"github.com/sebidchi/ft-quiz/internal/app/module/user-answer/domain"
	utils "github.com/sebidchi/ft-quiz/pkg"
)

func RandomUsersAnswers(username string) *domain.UserAnswers {
	return &domain.UserAnswers{
		UserId:  username,
		Answers: RandomUserAnswers(),
	}
}

func RandomUserAnswers() []*domain.UserAnswer {
	return []*domain.UserAnswer{
		utils.RandomElementFromSlice[*domain.UserAnswer]([]*domain.UserAnswer{RandomEqualUserAnswer(), RandomDiffUserAnswer()}),
		utils.RandomElementFromSlice[*domain.UserAnswer]([]*domain.UserAnswer{RandomEqualUserAnswer(), RandomDiffUserAnswer()}),
		utils.RandomElementFromSlice[*domain.UserAnswer]([]*domain.UserAnswer{RandomEqualUserAnswer(), RandomDiffUserAnswer()}),
		utils.RandomElementFromSlice[*domain.UserAnswer]([]*domain.UserAnswer{RandomEqualUserAnswer(), RandomDiffUserAnswer()}),
	}
}

func RandomEqualUserAnswer() *domain.UserAnswer {
	id := gofakeit.UUID()
	return domain.NewUserAnswer(
		id,
		id,
	)
}

func RandomDiffUserAnswer() *domain.UserAnswer {
	return domain.NewUserAnswer(
		gofakeit.UUID(),
		gofakeit.UUID(),
	)
}
