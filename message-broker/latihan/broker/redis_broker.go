package broker

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type RedisBroker struct {
	client *redis.Client
}

func NewRedisBroker(client *redis.Client) Broker {
	return &RedisBroker{client: client}
}

func (r *RedisBroker) Publish(ctx context.Context, topic string, data interface{}) error {
	return r.client.Publish(ctx, topic, data).Err()
}

func (r *RedisBroker) Subscribe(ctx context.Context, topic string, receiver Receiver) {
	subs := r.client.Subscribe(ctx, topic)
	for {
		message, err := subs.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}

		receiver(message)
	}
}
