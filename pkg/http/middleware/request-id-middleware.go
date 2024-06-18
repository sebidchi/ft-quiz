package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	utils "github.com/sebidchi/ft-quiz/pkg"
)

const HeaderXRequestID = "X-Request-OptionId"

type IdentifierGenerator func() string

type RequestIdMiddleware struct {
	IdGenerator IdentifierGenerator
}

func NewRequestIdMiddleware(generator IdentifierGenerator) *RequestIdMiddleware {
	return &RequestIdMiddleware{IdGenerator: generator}
}

func (m *RequestIdMiddleware) RequestIdentifier() gin.HandlerFunc {
	return func(c *gin.Context) {
		if m.IdGenerator == nil {
			m.IdGenerator = func() string {
				return utils.NewUlid().String()
			}
		}

		rid := c.Request.Header.Get(HeaderXRequestID)

		if rid == "" {
			rid = m.IdGenerator()
		}

		c.Request.Header.Set(HeaderXRequestID, rid)

		c.Request = c.Request.WithContext(context.WithValue(c, "correlation_id", c.Request.Header.Get("X-Request-ID")))

		c.Writer.Header().Set(HeaderXRequestID, rid)

		c.Next()
	}
}
