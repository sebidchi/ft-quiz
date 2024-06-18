package response

import (
	"net/http"
	"strconv"

	"github.com/google/jsonapi"
	utils "github.com/sebidchi/ft-quiz/pkg"
)

func NewInternalServerError() []*jsonapi.ErrorObject {
	return []*jsonapi.ErrorObject{{
		ID:     utils.NewUlid().String(),
		Code:   "internal_server_error",
		Title:  "Internal Server Error",
		Detail: "Internal Server Error",
		Status: strconv.Itoa(http.StatusInternalServerError),
	}}
}
