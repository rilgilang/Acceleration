package main

import (
	"context"
	"fmt"
	"htmltopdf/generator"
)

type User struct {
	Name string `json:"Name"`
}

func main() {
	ctx := context.Background()

	wow := generator.NewDocumentGenerator()

	err := wow.Resolve(ctx, User{Name: "Wow sangaddd"})

	if err != nil {
		fmt.Println("err --> ", err)
	}

}
