package queue_test

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"queue"
	"testing"
)

func TestNewQueueRedis(t *testing.T) {
	ctx := context.Background()

	rdbConnect := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       2,
	})

	rdb := queue.NewQueueRedis(rdbConnect)

	err := rdb.Push(ctx, &queue.Job{ID: "1", Message: "Job 1"})

	if err != nil {
		fmt.Println("err push --> ", err)
	}

	err = rdb.Push(ctx, &queue.Job{ID: "2", Message: "Job 2"})

	if err != nil {
		fmt.Println("err push --> ", err)
	}

	total := rdb.Count(ctx)
	fmt.Println("total --> ", total)

	task, err := rdb.Pop(ctx)
	if err != nil {
		fmt.Println("err pop --> ", err)
	}

	fmt.Println("Task --> ", task)

}
