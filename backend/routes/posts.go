package routes

import (
	"github.com/gofiber/fiber/v2"
)

func PostsRoutes(a *fiber.App) {
	route := a.Group("/api/v1/posts")

	route.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World from posts")
	})

	route.Post("/", func(c *fiber.Ctx) error {
		return c.SendString("Posted to posts:\n" + string(c.Body()))
	})

	route.Get("/:id", func(c *fiber.Ctx) error {
		return c.SendString("Get post by id")
	})

	route.Put("/:id", func(c *fiber.Ctx) error {
		return c.SendString("Update post by id:\n" + string(c.Body()))
	})
	
	route.Delete("/:id", func(c *fiber.Ctx) error {
		return c.SendString("Delete post by id")
	})

}