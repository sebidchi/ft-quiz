package response

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	utils "github.com/sebidchi/ft-quiz/pkg"
)

func NewNotFound(detail string) gin.H {
	return gin.H{
		"ID":     utils.NewUlid().String(),
		"Code":   "not_found",
		"Title":  "Not Found",
		"Detail": detail,
		"Status": strconv.Itoa(http.StatusNotFound),
	}
}
