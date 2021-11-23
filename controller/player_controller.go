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
	var response model.Response
	context.BindJSON(&request)
	if isValid, errorMessage := validateRequest(&request); !isValid {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": errorMessage})
		return
	}
	hasUnknown, prev, next := request.GetFirstUnkown()
	if !hasUnknown {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "El string enviado no tiene caracteres pendientes por descubrir"})
		return
	}
	response.Build(prev, next, request.Tries)
	context.IndentedJSON(http.StatusOK, response)
}

// validateRequest
func validateRequest(request *model.Request) (bool, string) {
	if request.WordToGuess == "" {
		return false, "No se recibió palabra para generar sugerencias."
	}
	return true, ""
}
