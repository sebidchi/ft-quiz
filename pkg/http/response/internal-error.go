package response

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	utils "github.com/sebidchi/ft-quiz/pkg"
)

func NewInternalServerError() gin.H {
	return gin.H{
		"ID":     utils.NewUlid().String(),
		"Code":   "internal_server_error",
		"Title":  "Internal Server Error",
		"Detail": "Internal Server Error",
		"Status": strconv.Itoa(http.StatusInternalServerError),
	}
}
