package routes

import (
	"github.com/gofiber/fiber/v2"
)

func PostsRoutes(a *fiber.App) {
	route := a.Group("/api/v1/posts")

	route.Get("/posts", func(c *fiber.Ctx) error {
		return c.SendString("Hello World from posts")
	})
}