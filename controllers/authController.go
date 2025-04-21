package controllers

import (
	"github.com/doneowth/models"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	user := models.User{
		FirstName: "don",
		LastName:  "liang",
	}

	return c.JSON(user)
}

func Hello(c *fiber.Ctx) error {
	return c.SendString("🔥 Auth 🔐")
}
