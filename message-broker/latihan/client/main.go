package main

import (
	"context"
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

	b.Publish(context.Background(), "topic-1", "ini topic 1 blok 😡")
}
