package ui

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sebidchi/ft-quiz/internal/app/module/user-answer/application"
	"github.com/sebidchi/ft-quiz/internal/pkg/domain/bus"
	"github.com/sebidchi/ft-quiz/internal/pkg/infrastructure/question"
	"github.com/sebidchi/ft-quiz/pkg/http/middleware"
	"github.com/sebidchi/ft-quiz/pkg/http/response"
)

const PostUserAnswersSchema = "post-user-answers.schema.json"

func HandlePostUserAnswers(commandBus bus.CommandBus, jar *middleware.JsonResponseMiddleware) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request := &question.AnswersPayload{}

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		answers := make(map[string]string)
		for _, answer := range request.Answers {
			answers[answer.QuestionID] = answer.Answer
		}
		storeCommand := application.NewStoreUserAnswersCommand(request.Username, answers)

		err := commandBus.Dispatch(ctx, storeCommand)

		switch err.(type) {
		case nil:
			jar.WriteResponse(ctx, nil, http.StatusNoContent)
			return
		default:
			jar.WriteErrorResponse(ctx, response.NewInternalServerError(), http.StatusInternalServerError, err)
			return
		}

	}
}
