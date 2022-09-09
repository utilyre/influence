package main

import (
	"fmt"
	"log"
	"os"
	"server/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	err := database.Connect()
	if err != nil {
		log.Fatalln(err)
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	addr := fmt.Sprintf("0.0.0.0:%s", os.Getenv("SERVER_PORT"))
	app.Listen(addr)
}
