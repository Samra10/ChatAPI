package handlers

import (
	"net/http"
	"strings"

	"ChatAPI/GoChat/configs"

	"github.com/gin-gonic/gin"
	"github.com/imroc/req"
)

type chatResponse struct {
	Number           int    `json:"number"`
	ApplicationToken string `json:"application_token"`
}

type appsResponse struct {
	Name             string `json:"name"`
	ApplicationToken string `json:"application_token"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	ChatsCount       int    `json:"chats_count"`
}

var appsChach = map[string]int{
	"123456789": 0,
	"987654321": 0,
}

func CreateChat(c *gin.Context) {

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

func GetChatData(applicationToken string) (appsResponse, error) {

	var response appsResponse

	//get Chats Data from Rails API
	resp, err := req.Get(strings.Replace(configs.AppAPIUrl+configs.ChatsRoute, "{application_token}", applicationToken, 1))
	if err != nil {
		return response, err
	}

	resp.ToJSON(&response)
	return response, nil

}
