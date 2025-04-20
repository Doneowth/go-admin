package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:my-secret-pw@/mydb"), &gorm.Config{})

	if err != nil {
		panic("hi db issue")
	}

	fmt.Print(db)
    // Initialize a new Fiber app
    app := fiber.New()

    // Define a route for the GET method on the root path '/'
    app.Get("/", func(c *fiber.Ctx) error {
        // Send a string response to the client
        return c.SendString("Hello, World 👋!")
    })

    // Start the server on port 8000
    log.Fatal(app.Listen(":8000"))
}