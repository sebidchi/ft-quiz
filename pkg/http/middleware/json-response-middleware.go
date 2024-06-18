package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sebidchi/ft-quiz/pkg/logger"
)

type JsonResponseMiddleware struct {
	lg logger.Logger
}

func NewJsonResponseMiddleware(logger logger.Logger) *JsonResponseMiddleware {
	return &JsonResponseMiddleware{lg: logger}
}

func (jar *JsonResponseMiddleware) WriteErrorResponse(ctx *gin.Context, errors interface{}, httpStatus int, previousError error) {
	jar.logError(previousError, ctx.GetHeader(HeaderXRequestID), httpStatus)

	ctx.JSON(httpStatus, errors)
}

func (jar *JsonResponseMiddleware) WriteResponse(ctx *gin.Context, payload interface{}, httpStatus int) {
	ctx.JSON(httpStatus, payload)
}

func (jar *JsonResponseMiddleware) logError(err error, correlationId string, statusCode int) {
	if err == nil {
		return
	}

	if statusCode >= http.StatusInternalServerError {
		logger.LogErrors(logger.Error, err, jar.lg, map[string]interface{}{"correlation_id": correlationId})
	} else {
		logger.LogErrors(logger.Warning, err, jar.lg, map[string]interface{}{"correlation_id": correlationId})
	}
}
