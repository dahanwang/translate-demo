package redis

import (
	"fmt"
	"github.com/go-redis/redis"
)

var Redisdb *redis.Client

func init() {
	Redisdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := Redisdb.Ping().Result()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Connect to redis Successfully")
}