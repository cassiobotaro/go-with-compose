package main

import (
	"fmt"
	redis "gopkg.in/redis.v4"
	"os"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}
