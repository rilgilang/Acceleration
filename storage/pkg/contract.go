package pkg

import "context"

type Storage interface {
	Put(ctx context.Context, bucket, outputPath string, file []byte) error
	Get(ctx context.Context, bucket, object string) ([]byte, error)
}
