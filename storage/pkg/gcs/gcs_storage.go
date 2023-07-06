package gcs

import (
	"bytes"
	"cloud.google.com/go/storage"
	"context"
	"errors"
	"fmt"
	"google.golang.org/api/option"
	"io"
	"storage/pkg"
)

type gcsStorage struct {
	client *storage.Client
}

// NewGCS creates a Google Cloud Storage Client
func NewGCS(ctx context.Context, cfgJsonFIle string) (pkg.Storage, error) {
	credOpt := option.WithCredentialsFile(cfgJsonFIle)

	client, err := storage.NewClient(ctx, credOpt)
	if err != nil {
		return nil, fmt.Errorf("storage.NewClient: %w", err)
	}

	return &gcsStorage{client}, nil
}

func (s *gcsStorage) Put(ctx context.Context, bucket, object string, file []byte) error {
	// Upload an object with storage.Writer.
	wc := s.client.Bucket(bucket).Object(object).NewWriter(ctx)

	if _, err := wc.Write(file); err != nil {
		return fmt.Errorf("storage.Writer.Write: %w", err)
	}

	if err := wc.Close(); err != nil {
		return fmt.Errorf("storage.Writer.Close: %w", err)
	}

	return nil
}

func (s *gcsStorage) Get(ctx context.Context, bucket, object string) ([]byte, error) {
	r, err := s.client.Bucket(bucket).Object(object).NewReader(ctx)
	if err != nil {
		if errors.Is(err, storage.ErrObjectNotExist) {
			return nil, storage.ErrObjectNotExist
		}
	}
	defer r.Close()

	var b bytes.Buffer
	if _, err := io.Copy(&b, r); err != nil {
		return nil, fmt.Errorf("failed to download bytes: %w", err)
	}

	return b.Bytes(), nil
}
