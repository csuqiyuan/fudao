package util

import (
	"fmt"
	"github.com/go-redis/redis"
)


func Get(key string) (data string){
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	data, err := client.Get(key).Result()
	if err == redis.Nil {
		fmt.Println(key+" does not exists")
	} else if err != nil {
		panic(err)
	}
	return
}
