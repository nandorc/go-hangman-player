package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PlayerService
// Retrieves character suggestions based on a word to guess and characters tried before
func PlayerService(context *gin.Context) {
	message := "Hola"
	context.IndentedJSON(http.StatusOK, message)
}
