package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type PostUserAnswers struct {
	IntegrationSuite
}

func (suite *PostUserAnswers) SetupSuite() {
	suite.IntegrationSuite.SetupSuite()
}

func (suite *PostUserAnswers) SetupTest() {
	suite.IntegrationSuite.SetupTest()
}

func TestPostUserAnswers(t *testing.T) {
	suite.Run(t, new(PostUserAnswers))
}

func (suite *PostUserAnswers) TestHandlePostUserAnswersHappyPath() {
	userName := "Rob pike"
	response := suite.executeJsonRequest(
		http.MethodPost,
		"/answers",
		[]byte(fmt.Sprintf(`{
			"username": "%s",
			"answers": [
				{
					"question_id": "option1",
					"answer": "option1"
				},
				{
					"question_id": "option5",
					"answer": "option5"
				},
				{
					"question_id": "option11",
					"answer": "option9"
				},
				{
					"question_id": "option13",
					"answer": "option13"
				},
				{
					"question_id": "option19",
					"answer": "option17"
				}
			]
}`, userName)),
		EmptyHeaders(),
	)

	suite.checkResponse(http.StatusNoContent,
		"",
		response,
	)
	suite.UserResultsShouldExist(userName)
}

func (suite *PostUserAnswers) TestHandlePostUserAnswersInvalidPayload() {
	userName := "Rob pike"
	suite.DbArranger.Arrange(suite.Ctx)
	response := suite.executeJsonRequest(
		http.MethodPost,
		"/answers",
		[]byte(fmt.Sprintf(`{
			"a": "%s",
			"answers": [
				{
					"question_id": "option1",
					"answer": "option1"
				},
				{
					"question_id": "option5",
					"answer": "option5"
				},
				{
					"question_id": "option11",
					"answer": "option9"
				},
				{
					"question_id": "option13",
					"answer": "option13"
				},
				{
					"question_id": "option19",
					"answer": "option17"
				}
			]
}`, userName)),
		EmptyHeaders(),
	)

	suite.checkResponseCode(http.StatusBadRequest, response.Code)
	suite.checkResponse(http.StatusBadRequest,
		"[{\"Code\":\"required\",\"Detail\":\"(root): username is required\",\"ID\":\"<<PRESENCE>>\",\"Meta\":{\"context\":\"(root)\",\"field\":\"(root)\",\"property\":\"username\"},\"Status\":\"400\",\"Title\":\"username is required\"},{\"Code\":\"additional_property_not_allowed\",\"Detail\":\"(root): Additional property a is not allowed\",\"ID\":\"<<PRESENCE>>\",\"Meta\":{\"context\":\"(root)\",\"field\":\"(root)\",\"property\":\"a\"},\"Status\":\"400\",\"Title\":\"Additional property a is not allowed\"}]",
		response,
	)
	suite.UserResultsShouldNotExist(userName)
}
