package redis

import (
	"ChatAPI/GoChat/configs"
	"context"
	"encoding/json"

	"github.com/bsm/redislock"
	"github.com/go-redis/redis/v8"
)

//to prevent create new client on each call
var redisClient *redis.Client
var redisLocker *redislock.Client // to prevent race condition

type sidekiqJob struct {
	Class string   `json:"class"`
	Args  []string `json:"args"`
	Retry bool     `json:"retry"`
	Queue string   `json:"queue"`
}

func GetRedisClient() *redis.Client {

	//check if client is nil
	if redisClient == nil {
		//create new client
		redisClient = redis.NewClient(&redis.Options{
			Addr:     configs.RedisUrl,
			Password: "", //no password set
			DB:       0,  // use default DB
		})
	}
	return redisClient
}

func GetRedisLocker() *redislock.Client {

	if redisClient == nil {
		GetRedisClient()
	}
	//defer redisClient.Close()
	redisLocker = redislock.New(redisClient)

	return redisLocker
}

func PushToRedis(ctx context.Context, queue string, class string, args ...string) error {
	job := sidekiqJob{
		Class: class,
		Args:  args,
		Queue: queue,
		Retry: true,
	}

	if redisClient == nil {
		GetRedisClient()
	}

	jobBytes, err := json.Marshal(job)
	if err != nil {
		return err
	}

	_, err = redisClient.RPush(ctx, "queue:"+queue, jobBytes).Result()
	return err
}
