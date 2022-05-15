package handlers

import (
	"ChatAPI/GoChat/configs"
	"ChatAPI/GoChat/redis"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/imroc/req"
)

type messageRequest struct {
	MessageBody string `json:"message_body"`
}

type messageResponse struct {
	Number           int64  `json:"number"`
	ChatNumber       int64  `json:"chat_number"`
	ApplicationToken string `json:"application_token"`
}

type ApiChatResponse struct {
	Number        int64  `json:"number"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	MessagesCount int64  `json:"meesages_count"`
}

func AddMessage(c *gin.Context) {
	// Read in request
	applicationToken := c.Param("application_token")
	chatNumber := c.Param("chat_number")

	messageBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	var req messageRequest
	err = json.Unmarshal(messageBody, &req)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	// Get next message number
	redisClient := redis.GetRedisClient()
	key := "MSG" + applicationToken + "_" + chatNumber

	exists, err := redisClient.Exists(key).Result()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	} else if exists == 0 {
		// Key not Exist, Get the messages counts from Rails API
		chatsResponse, err := RequestMessages(applicationToken, chatNumber)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
		redisClient.Set(key, chatsResponse.MessagesCount, 0)
	}
	//Increase the messages counts
	nextMessageNumber, err := redisClient.Incr(key).Result()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	// push update to workers
	err = redis.PushToRedis(configs.MessagesQueue, configs.MessagesWorker, applicationToken, chatNumber, strconv.FormatInt(nextMessageNumber, 10), req.MessageBody)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	// Send response
	chatNum, _ := strconv.ParseInt(chatNumber, 10, 64)
	response := messageResponse{
		Number:           nextMessageNumber,
		ChatNumber:       chatNum,
		ApplicationToken: applicationToken}

	c.IndentedJSON(http.StatusCreated, response)
}

func RequestMessages(applicationToken string, chatNumber string) (ApiChatResponse, error) {
	var response ApiChatResponse

	url := strings.Replace(configs.AppAPIUrl+configs.MessagesRoute, "{application_token}", applicationToken, 1)
	url = strings.Replace(url, "{chat_number}", chatNumber, 1)

	r, err := req.Get(url)
	if err != nil {
		return response, err
	}

	r.ToJSON(&response)
	return response, nil
}
