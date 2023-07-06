package broker

import "context"

type Broker interface {
	Publish(ctx context.Context, topic string, data interface{}) error
	Subscribe(ctx context.Context, topic string, receiver Receiver)
}

type Receiver func(interface{})
