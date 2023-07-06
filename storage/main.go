package main

import (
	"context"
	"fmt"
	"os"
	"storage/pkg/gcs"
	"storage/pkg/minio"
)

func main() {
	ctx := context.Background()

	// Open the file for reading
	file, err := os.ReadFile("./example.pdf")
	if err != nil {
		fmt.Println("error --> ", err)
	}

	gcsClient, err := gcs.NewGCS(ctx, jsonFile)
	minioClient, err := minio.NewMinio(ctx)

	if err != nil {
		panic(fmt.Sprintf(`error connecting minio --> %v`, err))
	}

	err = minioClient.Put(ctx, "waduh", "/yeaa/boi.pdf", file)

	if err != nil {
		fmt.Println("error --> ", err)
	}

	ajg, err := minioClient.Get(ctx, "waduh", "/yeaa/boi.pdf")

	if err != nil {
		fmt.Println("error --> ", err)
	}

	fmt.Println("ajg -->", ajg)
}
