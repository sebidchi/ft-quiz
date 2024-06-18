package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type GetQuizSuite struct {
	IntegrationSuite
}

func (suite *GetQuizSuite) SetupSuite() {
	suite.IntegrationSuite.SetupSuite()
}

func (suite *GetQuizSuite) SetupTest() {
	suite.IntegrationSuite.SetupTest()
}

func TestGetQuizSuite(t *testing.T) {
	suite.Run(t, new(GetQuizSuite))
}

func (suite *GetQuizSuite) TestHandleGetQuizHappyPath() {
	response := suite.executeJsonRequest(
		http.MethodGet,
		"/quiz",
		nil,
		EmptyHeaders(),
	)

	suite.checkResponse(http.StatusOK,
		"[{\"question\":\"question1\",\"questions\":\"What is the capital of France?\",\"options\":[{\"id\":\"option1\",\"text\":\"Paris\"},{\"id\":\"option2\",\"text\":\"Berlin\"},{\"id\":\"option3\",\"text\":\"London\"},{\"id\":\"option4\",\"text\":\"Madrid\"}],\"answer\":\"option1\"},{\"question\":\"question2\",\"questions\":\"Who wrote the novel '1984'?\",\"options\":[{\"id\":\"option5\",\"text\":\"George Orwell\"},{\"id\":\"option6\",\"text\":\"Aldous Huxley\"},{\"id\":\"option7\",\"text\":\"Ernest Hemingway\"},{\"id\":\"option8\",\"text\":\"Mark Twain\"}],\"answer\":\"option5\"},{\"question\":\"question3\",\"questions\":\"What is the square root of 81?\",\"options\":[{\"id\":\"option9\",\"text\":\"7\"},{\"id\":\"option10\",\"text\":\"8\"},{\"id\":\"option11\",\"text\":\"9\"},{\"id\":\"option12\",\"text\":\"10\"}],\"answer\":\"option11\"},{\"question\":\"question4\",\"questions\":\"What is the chemical symbol for Hydrogen?\",\"options\":[{\"id\":\"option13\",\"text\":\"H\"},{\"id\":\"option14\",\"text\":\"He\"},{\"id\":\"option15\",\"text\":\"Hy\"},{\"id\":\"option16\",\"text\":\"Ho\"}],\"answer\":\"option13\"},{\"question\":\"question5\",\"questions\":\"Who painted the Mona Lisa?\",\"options\":[{\"id\":\"option17\",\"text\":\"Vincent van Gogh\"},{\"id\":\"option18\",\"text\":\"Pablo Picasso\"},{\"id\":\"option19\",\"text\":\"Leonardo da Vinci\"},{\"id\":\"option20\",\"text\":\"Claude Monet\"}],\"answer\":\"option19\"}]",
		response,
	)
}
