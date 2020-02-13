package main

import (
	"github.com/astaxie/beego"
	_ "translate-demo/routers"
	"github.com/go-redis/redis"
)

var redisdb *redis.Client

func main() {

	//redis
	/**
	redisdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := redisdb.Ping().Result()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Connect to redis Successfully")
	 */

	// beego
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
