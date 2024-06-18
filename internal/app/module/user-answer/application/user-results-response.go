package application

import "github.com/sebidchi/ft-quiz/internal/app/module/user-answer/domain"

type UserResultsResponse struct {
	UserId     string
	Total      int
	Percentage float64
	BetterThan float64
	TotalUsers int
}

func NewUserResultsResponse(userResults *domain.UserResults) *UserResultsResponse {
	return &UserResultsResponse{
		UserId:     userResults.UserId,
		Total:      userResults.Total,
		Percentage: userResults.Percentage,
		BetterThan: userResults.BetterThan,
		TotalUsers: userResults.TotalUsers,
	}
}
