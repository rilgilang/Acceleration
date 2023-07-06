package queue_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"queue"
	"testing"
)

func TestNewQueue(t *testing.T) {
	db, err := sql.Open("mysql",
		"root:@tcp(127.0.0.1:3306)/learn_queue_mysql")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := queue.NewQueue(db)

	ctx := context.Background()

	getOne, err := sql.Pop(ctx)

	if err != nil {
		panic(err)
	}

	if getOne == nil {
		fmt.Println("no data found")
	} else {
		fmt.Println("resp --> ", getOne)
	}

}
