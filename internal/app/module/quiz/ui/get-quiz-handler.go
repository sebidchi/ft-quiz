package ui

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sebidchi/ft-quiz/internal/app/module/quiz/application"
	"github.com/sebidchi/ft-quiz/internal/app/module/quiz/domain"
	"github.com/sebidchi/ft-quiz/internal/pkg/domain/bus"
	"github.com/sebidchi/ft-quiz/internal/pkg/infrastructure/question"
	"github.com/sebidchi/ft-quiz/pkg/http/middleware"
	"github.com/sebidchi/ft-quiz/pkg/http/response"
)

func HandleGetQuiz(bus bus.QueryBus, jar *middleware.JsonResponseMiddleware) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		fetchQuizQuery := application.NewFetchQuizQuery("default")

		resp, err := bus.Ask(ctx, fetchQuizQuery)

		switch err.(type) {
		case nil:
			jar.WriteResponse(ctx, buildResponse(resp), http.StatusOK)
			return
		case *domain.QuizNotFound:
			jar.WriteErrorResponse(ctx, response.NewNotFound(err.Error()), http.StatusNotFound, err)
			return
		default:
			jar.WriteErrorResponse(ctx, response.NewInternalServerError(), http.StatusInternalServerError, err)
			return
		}

	}
}

func buildResponse(resp interface{}) []*question.Question {
	response := resp.(*application.QuizResponse)
	questions := make([]*question.Question, 0)
	for i := range response.Questions {
		q := response.Questions[i]
		options := make([]*question.Option, 0)
		for _, o := range q.Options {
			option := &question.Option{
				OptionId: o.Id,
				Text:     o.Text,
			}
			options = append(options, option)
		}

		questions = append(questions, &question.Question{
			ID:             q.Id,
			Question:       q.QuestionText,
			Options:        options,
			AnswerOptionId: q.Answer,
		})
	}

	return questions
}
