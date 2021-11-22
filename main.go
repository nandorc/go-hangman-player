package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/hangman", PlayerService)
	router.Run(":8090")
}

func PlayerService(context *gin.Context) {
	message := "Hola"
	context.IndentedJSON(http.StatusOK, message)
}
