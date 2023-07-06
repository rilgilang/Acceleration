package queue

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
)

type queueRedis struct {
	rd *redis.Client
}

func NewQueueRedis(rd *redis.Client) Queue {
	return &queueRedis{rd: rd}
}

func (qr *queueRedis) Count(ctx context.Context) int64 {
	total := qr.rd.LLen(ctx, "job")
	return total.Val()
}

func (qr *queueRedis) Push(ctx context.Context, job *Job) error {
	b, _ := json.Marshal(&job)
	err := qr.rd.RPush(ctx, "job", b).Err()
	return err
}

func (qr *queueRedis) Pop(ctx context.Context) (*Job, error) {
	var task = Job{}
	val, err := qr.rd.LPop(ctx, "job").Result()

	_ = json.Unmarshal([]byte(val), &task)

	return &task, err
}
