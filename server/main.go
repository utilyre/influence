package main

import (
	"fmt"
	"log"
	"os"
	"server/database"
	"server/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	err := database.Connect()
	if err != nil {
		log.Fatalln(err)
	}

	app := fiber.New()
  setupRoutes(app)

	addr := fmt.Sprintf("0.0.0.0:%s", os.Getenv("SERVER_PORT"))
	app.Listen(addr)
}

func setupRoutes(app *fiber.App) {
	api := app.Group("/api")

	routes.SetupUserRoutes(api)
}
