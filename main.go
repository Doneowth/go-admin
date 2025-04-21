package main

import (
	"log"

	"github.com/doneowth/database"
	"github.com/doneowth/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	// Initialize a new Fiber app
	app := fiber.New()

	routes.Setup(app)
	app.Listen(":8000") //nolint
	log.Println("all set")
}
