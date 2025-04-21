package routes

import (
	_ "log"

	"github.com/doneowth/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Hello)

	setPost(app)
}

func setPost(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
}
