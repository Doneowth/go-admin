package controllers

import (
	"log"

	"github.com/doneowth/database"
	"github.com/doneowth/interfaces"
	"github.com/doneowth/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	Service interfaces.UserService
}

func NewUserController(service interfaces.UserService) *UserController {
	return &UserController{Service: service}
}

func (c *UserController) CreateUser(ctx *fiber.Ctx) error {
	var data map[string]string
	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		ctx.Status(400)

		return ctx.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}
	
	pw, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	
	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:  data["email"],
		Password:  pw,
	}

	log.Printf("user: %v\n", user)

	createdUser, err := c.Service.CreateUser(&user)
	if err != nil {
		ctx.Status(500)

		return ctx.JSON(fiber.Map{
			"message": "server issue",
		})
	}
	return ctx.JSON(*createdUser)
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	var data map[string]string
	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		ctx.Status(400)

		return ctx.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}
	
	pw, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	
	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:  data["email"],
		Password:  pw,
	}

	database.DB.Create(&user)


	return ctx.JSON(user)
}