package main

import (
	"hangman-player/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/hangman", controller.PlayerService)
	router.Run(":8090")
}
