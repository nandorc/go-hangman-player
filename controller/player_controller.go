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
	if isValid, errorMessage := validateRequest(&request); !isValid {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": errorMessage})
		return
	}
	context.IndentedJSON(http.StatusOK, request)
}

// validateRequest
func validateRequest(request *model.Request) (bool, string) {
	if request.WordToGuess == "" {
		return false, "No se recibió palabra para generar sugerencias."
	}
	return true, ""
}
