package controller

import (
	"hangman-player/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PlayerService
// Retrieves character suggestions based on a word to guess and characters tried before
func PlayerService(context *gin.Context) {
	var request model.Request
	context.BindJSON(&request)
	context.IndentedJSON(http.StatusOK, request)
}
