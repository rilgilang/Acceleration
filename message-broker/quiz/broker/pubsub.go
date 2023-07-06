package broker

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type pubsubBroker struct {
	client *pubsub.Client
}

func NewPubsubBroker(client *pubsub.Client) Broker {
	return &pubsubBroker{
		client: client,
	}
}

func (ps *pubsubBroker) Publish(ctx context.Context, topic string, data interface{}) error {
	t := ps.client.Topic(topic)

	payload, _ := json.Marshal(data)

	result := t.Publish(ctx, &pubsub.Message{Data: payload, PublishTime: time.Now()})

	_, err := result.Get(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (ps *pubsubBroker) Subscribe(ctx context.Context, topic string, receiver Receiver) {
	subscriptionId := fmt.Sprintf(`%s-sub`, topic)

	subscribe := ps.client.Subscription(subscriptionId)

	err := subscribe.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {
		msg.Ack()
		receiver(string(msg.Data))
	})

	if err != nil {
		panic(err)
	}
}
