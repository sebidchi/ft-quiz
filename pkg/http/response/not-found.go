package response

import (
	"net/http"
	"strconv"

	"github.com/google/jsonapi"
	utils "github.com/sebidchi/ft-quiz/pkg"
)

func NewNotFound(detail string) []*jsonapi.ErrorObject {
	return []*jsonapi.ErrorObject{{
		ID:     utils.NewUlid().String(),
		Code:   "not_found",
		Title:  "Not Found",
		Detail: detail,
		Status: strconv.Itoa(http.StatusNotFound),
	}}
}
