package ui

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sebidchi/ft-quiz/internal/app/module/user-answer/application"
	"github.com/sebidchi/ft-quiz/internal/app/module/user-answer/domain"
	"github.com/sebidchi/ft-quiz/internal/pkg/domain/bus"
	"github.com/sebidchi/ft-quiz/internal/pkg/infrastructure/question"
	"github.com/sebidchi/ft-quiz/pkg/http/middleware"
	"github.com/sebidchi/ft-quiz/pkg/http/response"
)

func HandleGetUserResults(bus bus.QueryBus, jar *middleware.JsonResponseMiddleware) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userName := ctx.Param("username")

		query := application.NewGetUserResultsQuery(userName)
		resp, err := bus.Ask(ctx, query)

		switch err.(type) {
		case nil:
			jar.WriteResponse(ctx, buildResponse(resp), http.StatusOK)
			return
		case *domain.UsersResultsNotFound:
			jar.WriteErrorResponse(ctx, response.NewNotFound(err.Error()), http.StatusNotFound, err)
			return
		default:
			jar.WriteErrorResponse(ctx, response.NewInternalServerError(), http.StatusInternalServerError, err)
			return
		}

	}
}

func buildResponse(resp interface{}) question.UserResults {
	var userResults question.UserResults
	res, ok := resp.(*application.UserResultsResponse)
	if !ok {
		return userResults
	}
	return question.UserResults{
		UserId:     res.UserId,
		Total:      res.Total,
		Percentage: res.Percentage,
		BetterThan: res.BetterThan,
		TotalUsers: res.TotalUsers,
	}
}
