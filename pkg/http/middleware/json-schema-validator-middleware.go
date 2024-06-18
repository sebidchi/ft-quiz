package middleware

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	utils "github.com/sebidchi/ft-quiz/pkg"
	"github.com/sebidchi/ft-quiz/pkg/http/response"
	"github.com/xeipuuv/gojsonschema"
)

type RequestValidatorMiddleware struct {
	responseMiddleware *JsonResponseMiddleware
}

func NewRequestValidatorMiddleware(responseMiddleware *JsonResponseMiddleware) *RequestValidatorMiddleware {
	return &RequestValidatorMiddleware{
		responseMiddleware: responseMiddleware,
	}
}

func (jsv *RequestValidatorMiddleware) Validate(schemaNameLocation string) gin.HandlerFunc {
	return func(c *gin.Context) {
		absPath, _ := filepath.Abs(schemaNameLocation)
		schemaLoader := gojsonschema.NewReferenceLoader(fmt.Sprintf("file://%s", absPath))

		newRequest := utils.CloneRequest(c.Request)
		bodyBytes, err := io.ReadAll(newRequest.Body)

		if err != nil {
			jsv.responseMiddleware.WriteErrorResponse(c, response.NewBadRequest("Invalid payload received"), http.StatusBadRequest, err)
			c.Abort()
			return
		}

		documentLoader := gojsonschema.NewBytesLoader(bodyBytes)

		result, err := gojsonschema.Validate(schemaLoader, documentLoader)
		if err != nil {
			jsv.responseMiddleware.WriteErrorResponse(c, response.NewBadRequest("Invalid payload received"), http.StatusBadRequest, err)
			c.Abort()
			return
		}

		if !result.Valid() {
			errors := result.Errors()
			var validationErrors []gin.H
			for i := range errors {
				desc := errors[i]
				var details map[string]interface{}
				details = desc.Details()
				validationErrors = append(validationErrors, gin.H{
					"ID":     utils.NewUlid().String(),
					"Code":   desc.Type(),
					"Title":  desc.Description(),
					"Detail": desc.String(),
					"Status": "400",
					"Meta":   &details,
				})
			}

			jsv.responseMiddleware.WriteErrorResponse(c, validationErrors, http.StatusBadRequest, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
