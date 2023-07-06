package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Transfer struct {
	IdempotenceId int    `json:"idempotence_id"`
	To            string `json:"to"`
	Amount        string `json:"amount"`
}

var TransferList = []Transfer{}

func main() {
	// Fiber instance
	app := fiber.New()

	// Routes
	app.Post("/transfer", TransferHandler)

	// Start server
	log.Fatal(app.Listen(":6969"))
}

// Handler
func TransferHandler(c *fiber.Ctx) error {
	var params = Transfer{}

	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	for _, tr := range TransferList {
		if tr.IdempotenceId == params.IdempotenceId {
			return c.SendString(fmt.Sprintf(`You've already transfer to %s`, params.To))
		}
	}

	TransferList = append(TransferList, params)

	time.Sleep(time.Duration(rand.Intn(9-1)+1) * time.Second)

	return c.SendString(fmt.Sprintf(`You transfer to %s %v`, params.To, params.Amount))
}
