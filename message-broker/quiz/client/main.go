package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"google.golang.org/api/option"
	"pubsub/broker"
)

func main() {
	ctx := context.Background()
	projectId := "learn-golang-pubsub"
	topic := "sample-topic"

	cfg := option.WithCredentialsFile("../../learn-golang-pubsub-565b1c1cbb8c.json")

	pubsubClient, err := pubsub.NewClient(ctx, projectId, cfg)
	if err != nil {
		panic(err)
	}

	defer pubsubClient.Close()

	pubsubBroker := broker.NewPubsubBroker(pubsubClient)

	err = pubsubBroker.Publish(ctx, topic, "bejir")

	if err != nil {
		fmt.Println("error --> ", err)
	}

}
