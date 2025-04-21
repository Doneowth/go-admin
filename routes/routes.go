package routes

import (
	_ "log"

	"github.com/doneowth/controllers"
	"github.com/doneowth/database"
	"github.com/doneowth/services"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	setGetIndex(app)
	setPost(app)
}

func setGetIndex(app *fiber.App) {
	app.Get("/", controllers.Hello)
}

func setPost(app *fiber.App) {
	userService := services.NewUserService(database.DB)
	userController := controllers.NewUserController(userService)
	app.Post("/api/register", userController.CreateUser)
}
