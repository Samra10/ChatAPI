package handlers

import (
	"ChatAPI/GoChat/configs"
	"ChatAPI/GoChat/redis"
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/imroc/req"
)

type chatResponse struct {
	Number           int64  `json:"number"`
	ApplicationToken string `json:"application_token"`
}

type ApiAppsResponse struct {
	Name             string `json:"name"`
	ApplicationToken string `json:"application_token"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	ChatsCount       int64  `json:"chats_count"`
}

func CreateChat(c *gin.Context) {

	applicationToken := c.Param("application_token")
	// Get redis client from redis package
	redisClient := redis.GetRedisClient()
	// Get the redis lock to prevent race condition
	redisLocker := redis.GetRedisLocker()
	ctx := context.Background()

	//create the key
	key := "CHATS" + applicationToken

	// Try to obtain lock.
	lock, err := redisLocker.Obtain(ctx, key, 100*time.Millisecond, nil)
	if err != nil {

		defer lock.Release(ctx)
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	//check if the key exist
	exist, err := redisClient.Exists(ctx, key).Result()
	if err != nil {
		defer lock.Release(ctx)
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	} else if exist == 0 {
		// Key not Exist, Get the chat counts from Rails API
		resp, err := GetChatData(applicationToken)
		if err != nil {
			defer lock.Release(ctx)
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}
		//set the key with chat counts
		redisClient.Set(ctx, key, resp.ChatsCount, 0)
	}
	//Increase the Chact counts
	nextChatNumber, err := redisClient.Incr(ctx, key).Result()
	defer lock.Release(ctx)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	// push update to workers
	err = redis.PushToRedis(ctx, configs.ChatQueue, configs.ChatWorker, applicationToken, strconv.FormatInt(nextChatNumber, 10))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	// send response
	newChat := chatResponse{
		Number:           nextChatNumber,
		ApplicationToken: applicationToken,
	}
	c.IndentedJSON(http.StatusCreated, newChat)

}

func GetChatData(applicationToken string) (ApiAppsResponse, error) {

	var response ApiAppsResponse

	//get Chats Data from Rails API
	resp, err := req.Get(strings.Replace(configs.AppAPIUrl+configs.ChatsRoute, "{application_token}", applicationToken, 1))
	if err != nil {
		return response, err
	}

	resp.ToJSON(&response)
	return response, nil

}
