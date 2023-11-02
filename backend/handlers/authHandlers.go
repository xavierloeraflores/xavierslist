package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterUser(c *fiber.Ctx) error {
	return c.SendString("Register User")
}

func LoginUser(c *fiber.Ctx) error {
	return c.SendString("Login User")
}
