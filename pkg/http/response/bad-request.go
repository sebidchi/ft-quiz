package response

import (
	"net/http"
	"strconv"

	"github.com/google/jsonapi"
	utils "github.com/sebidchi/ft-quiz/pkg"
)

func NewBadRequest(detail string) []*jsonapi.ErrorObject {
	return []*jsonapi.ErrorObject{{
		ID:     utils.NewUlid().String(),
		Code:   "bad_request",
		Title:  "Bad Request",
		Detail: detail,
		Status: strconv.Itoa(http.StatusBadRequest),
	}}
}
