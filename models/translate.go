package models

import (
	"fmt"
	"translate-demo/redis"
)

func HGetTranslate(md5 string, tolg string) string{
	result,err := redis.Redisdb.HGet(md5, tolg).Result()

	if err != nil {
		fmt.Println("redis get error!! err = " + err.Error())
		return ""
	}

	return result
}

func HSetTranslat(md5 string, tolg string, translate string) {
	err := redis.Redisdb.HSet(md5,tolg,translate).Err()

	if err != nil {
		fmt.Println("redis set error!! err = " +  err.Error())
	}
}
