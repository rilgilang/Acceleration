package minio

import (
	"bytes"
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io/ioutil"
	"log"
	"storage/pkg"
)

type minioStorage struct {
	client *minio.Client
}

func NewMinio(ctx context.Context) (pkg.Storage, error) {

	endpoint := "localhost:9000"
	accessKeyID := "crNGKhDQBxXpbVF34VjL"
	secretAccessKey := "oT3nXgKD1KK6e6JCKJPnMx56yhoe2leLnL4PtnhY"
	bucketName := "waduh"

	// Initialize minio client object.
	minioClient, errInit := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if errInit != nil {
		return nil, errInit
	}

	err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: "us-east-1"})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			return nil, err
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	return &minioStorage{minioClient}, nil
}

func (m *minioStorage) Put(ctx context.Context, bucket, outputPath string, file []byte) error {

	reader := bytes.NewReader(file)

	// Set the content type of the file
	contentType := "application/octet-stream"

	// Upload the file to the bucket
	_, err := m.client.PutObject(ctx, bucket, outputPath, reader, -1,
		minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return err
	}

	return nil
}

func (m *minioStorage) Get(ctx context.Context, bucket, object string) ([]byte, error) {
	// Download the file from the bucket
	err := m.client.FGetObject(ctx, bucket, object, "./download/test_download.pdf", minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	fileByte, err := ioutil.ReadFile("./download/test_download.pdf")
	if err != nil {
		return nil, err
	}
	return fileByte, nil
}
