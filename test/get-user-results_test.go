package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/sebidchi/ft-quiz/internal/app/module/user-answer/domain"
	"github.com/stretchr/testify/suite"
)

type GetUserResults struct {
	IntegrationSuite
}

func (suite *GetUserResults) SetupSuite() {
	suite.IntegrationSuite.SetupSuite()
}

func (suite *GetUserResults) SetupTest() {
	suite.IntegrationSuite.SetupTest()
}

func TestGetUserResults(t *testing.T) {
	suite.Run(t, new(GetUserResults))
}

func (suite *GetUserResults) TestHandleGetUserResultsHappyPath() {
	userName := "Sebastian Parra"
	userAnswers := RandomUsersAnswers(userName)
	usersResults, err := domain.UsersResults{userName: domain.Result{
		Points:         userAnswers.Score(),
		TotalQuestions: len(userAnswers.Answers),
	}}.UserResults(userName)
	suite.NoError(err)
	suite.GivenUserAnswers(userAnswers)
	response := suite.executeJsonRequest(
		http.MethodGet,
		fmt.Sprintf("/answers/%s", userName),
		nil,
		EmptyHeaders(),
	)

	suite.checkResponse(http.StatusOK,
		fmt.Sprintf(
			"{\"user_id\":\"%s\",\"total\":%d,\"percentage\":%f,\"better_than\":%f,\"total_users\":%d}",
			userName,
			usersResults.Total,
			usersResults.Percentage,
			usersResults.BetterThan,
			usersResults.TotalUsers,
		),
		response,
	)
}

func (suite *GetUserResults) TestHandleGetUserResultsNotFound() {
	userName := "Sebastian Parra"
	notFoundName := "not-found"
	userAnswers := RandomUsersAnswers(userName)
	suite.GivenUserAnswers(userAnswers)

	response := suite.executeJsonRequest(
		http.MethodGet,
		fmt.Sprintf("/answers/%s", notFoundName),
		nil,
		EmptyHeaders(),
	)

	suite.checkResponse(
		http.StatusNotFound,
		"{\"Code\":\"not_found\",\"Detail\":\"user has no results yet\",\"ID\":\"<<PRESENCE>>\",\"Status\":\"404\",\"Title\":\"Not Found\"}\n",
		response,
	)
}
