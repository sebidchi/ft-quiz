package test

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/kinbiko/jsonassert"
	"github.com/sebidchi/ft-quiz/cmd/di"
	"github.com/sebidchi/ft-quiz/internal/app/module/user-answer/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type IntegrationSuite struct {
	suite.Suite
	Ctx        context.Context
	DI         *di.FTQuizDI
	DbArranger *InMemoryDbArranger
}

func (suite *IntegrationSuite) SetupSuite() {
	if suite.DI == nil {
		suite.DI = di.InitWithEnvFile("../.env", "../.env.test")
	}
	suite.DbArranger = NewInMemoryDbArranger(suite.DI)
	suite.Ctx = context.Background()
}

func (suite *IntegrationSuite) SetupTest() {
}

func (suite *IntegrationSuite) TearDownTest() {
}

func (suite *IntegrationSuite) TearDownSuite() {
}

func (suite *IntegrationSuite) GivenUserAnswers(answers *domain.UserAnswers) {
	suite.NoError(suite.DI.UserAnswersServices.UserAnswerRepository.SaveUserAnswers(suite.Ctx, answers))
}

func (suite *IntegrationSuite) UserResultsShouldExist(userId string) {
	res, err := suite.DI.UserAnswersServices.UserAnswerRepository.GetUsersResults(suite.Ctx)
	suite.NoError(err)
	results, err := res.UserResults(userId)
	suite.NoError(err)
	suite.NotNil(results)
}

func (suite *IntegrationSuite) UserResultsShouldNotExist(userId string) {
	res, err := suite.DI.UserAnswersServices.UserAnswerRepository.GetUsersResults(suite.Ctx)
	suite.NoError(err)
	results, err := res.UserResults(userId)
	suite.Error(err)
	suite.Nil(results)
}

func (suite *IntegrationSuite) executeJsonRequest(verb string, path string, body []byte, headers map[string]string) *httptest.ResponseRecorder {
	req, err := http.NewRequest(verb, path, bytes.NewBuffer(body))

	if len(headers) != 0 {
		for headerName, value := range headers {
			req.Header.Set(headerName, value)
		}
	}

	assert.NoError(suite.T(), err)

	req.Header.Set("Content-Type", "application/json")
	return suite.executeRequest(req)
}

func (suite *IntegrationSuite) executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()

	suite.DI.Services.Router.ServeHTTP(rr, req)

	return rr
}

func (suite *IntegrationSuite) checkResponse(expectedStatusCode int, expectedResponse string, response *httptest.ResponseRecorder, formats ...interface{}) {
	ja := jsonassert.New(suite.T())
	suite.checkResponseCode(expectedStatusCode, response.Code)

	receivedResponse := response.Body.String()
	fmt.Println(receivedResponse)
	fmt.Println(expectedResponse)
	if receivedResponse == "" {
		assert.Equal(suite.T(), expectedResponse, receivedResponse)
		return
	}
	if formats != nil {
		ja.Assertf(receivedResponse, expectedResponse, formats)
	} else {
		ja.Assertf(receivedResponse, expectedResponse)
	}
}

func (suite *IntegrationSuite) checkResponseCode(expected, actual int) {
	if expected != actual {
		suite.T().Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
