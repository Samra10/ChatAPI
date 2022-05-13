package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type chatResponse struct {
	Number           int    `json:"number"`
	ApplicationToken string `json:"application_toke"`
}

var appsChach = map[string]int{
	"123456789": 0,
	"987654321": 0,
}

func main() {
	router := gin.Default()

	router.GET("/applications/:application_token/chats", createChat)

	router.Run("localhost:8080")
}

func createChat(c *gin.Context) {

	var newChat chatResponse

	applicationToken := c.Param("application_token")

	//check if the applicationToken on the cach
	if val, ok := appsChach[applicationToken]; ok {

		//update number of chats
		appsChach[applicationToken] = val + 1
		val = val + 1

		newChat.ApplicationToken = applicationToken
		newChat.Number = val

		//return new chat
		c.IndentedJSON(http.StatusCreated, newChat)
		return
	}
	//return if application token not found
	c.IndentedJSON(http.StatusNotFound, gin.H{"Error": "ApplicationToken not found"})
}
