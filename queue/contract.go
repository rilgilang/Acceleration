package queue

import "context"

type Job struct {
	ID      string
	Message string
}

type Queue interface {
	Count(ctx context.Context) int64
	Push(ctx context.Context, job *Job) error
	Pop(ctx context.Context) (*Job, error)
}
