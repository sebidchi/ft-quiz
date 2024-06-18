package di

import (
	"github.com/sebidchi/ft-quiz/pkg/http/middleware"
	"github.com/sebidchi/ft-quiz/pkg/logger"
)

type httpServices struct {
	JsonResponseMiddleware *middleware.JsonResponseMiddleware
	RequestPanicMiddleware *middleware.RequestPanicMiddleware
	JsonSchemaValidator    *middleware.RequestValidatorMiddleware
	RequestIdMiddleware    *middleware.RequestIdMiddleware
}

func initHttpServices(logger logger.Logger) *httpServices {
	jsonResponseMiddleware := middleware.NewJsonResponseMiddleware(logger)
	return &httpServices{
		JsonResponseMiddleware: jsonResponseMiddleware,
		RequestPanicMiddleware: middleware.NewRequestPanicMiddleware(logger),
		JsonSchemaValidator:    middleware.NewRequestValidatorMiddleware(jsonResponseMiddleware),
		RequestIdMiddleware:    middleware.NewRequestIdMiddleware(nil),
	}
}
