package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sebidchi/ft-quiz/pkg/logger"
	"go.uber.org/zap"
)

type RequestPanicMiddleware struct {
	l logger.Logger
}

func NewRequestPanicMiddleware(l logger.Logger) *RequestPanicMiddleware {
	return &RequestPanicMiddleware{l: l}
}

func (rp *RequestPanicMiddleware) RequestPanicHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				rp.l.Error("Unhandled Error", zap.Any("error", err))
				jsonBody, _ := json.Marshal(map[string]interface{}{
					"errors": []map[string]string{
						{
							"title": "Internal Server Error",
						},
					},
				})

				c.Writer.Header().Set("Content-Type", "application/json")
				c.Writer.WriteHeader(http.StatusInternalServerError)
				if _, err := c.Writer.Write(jsonBody); err != nil {
					rp.l.Error("Could not write the response of the panic error", zap.Error(err))
				}
			}
		}()

		c.Next()
	}
}
