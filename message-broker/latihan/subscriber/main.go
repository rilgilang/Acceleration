package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"message-broker/broker"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	b := broker.NewRedisBroker(rdb)

	b.Subscribe(context.Background(), "topic-1", func(data interface{}) {
		fmt.Println(fmt.Sprintf(`ini message --> %v`, data))
	})
}
