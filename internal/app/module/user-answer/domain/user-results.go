package domain

import (
	"math"

	"github.com/sebidchi/ft-quiz/internal/pkg/domain"
)

type Result struct {
	Points         int
	TotalQuestions int
}
type UsersResults map[string]Result

type UserResults struct {
	UserId     string
	Total      int
	Percentage float64
	BetterThan float64
	TotalUsers int
}

func (u UsersResults) UserResults(userId string) (*UserResults, error) {
	userResults, ok := u[userId]
	if !ok {
		return nil, NewUsersResultsNotFound(userId)
	}
	totalUsers := len(u)
	betterThan := 0

	for uId, user := range u {
		if uId != userId && user.Points < userResults.Points {
			betterThan++
		}
	}

	comparison := float64(betterThan) / float64(totalUsers-1) * 100
	if len(u) == 1 {
		comparison = 0
	}
	percentage := float64(userResults.Points) / float64(userResults.TotalQuestions) * 100

	return &UserResults{
		UserId:     userId,
		Total:      userResults.Points,
		Percentage: math.Round(percentage*100) / 100,
		BetterThan: math.Round(comparison*100) / 100,
		TotalUsers: totalUsers,
	}, nil
}

type UsersResultsNotFound struct {
	domain.DomainError
	extraItems map[string]interface{}
}

func NewUsersResultsNotFound(userId string) *UsersResultsNotFound {
	return &UsersResultsNotFound{extraItems: map[string]interface{}{"userId": userId}}
}

func (u *UsersResultsNotFound) ExtraItems() map[string]interface{} {
	return u.extraItems
}

func (u *UsersResultsNotFound) Error() string {
	return "user has no results yet"
}
