package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetPosts(c *fiber.Ctx) error {
	return c.SendString("Hello World from posts")
}
