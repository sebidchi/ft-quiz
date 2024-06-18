package response

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	utils "github.com/sebidchi/ft-quiz/pkg"
)

func NewBadRequest(detail string) gin.H {
	return gin.H{
		"id":     utils.NewUlid().String(),
		"code":   "bad_request",
		"title":  "Bad Request",
		"detail": detail,
		"status": strconv.Itoa(http.StatusBadRequest),
	}
}
