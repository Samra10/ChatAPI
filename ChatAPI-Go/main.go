package main

import (
	"ChatAPI/GoChat/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/applications/:application_token/chats", handlers.CreateChat)

	router.Run("localhost:8080")
}
