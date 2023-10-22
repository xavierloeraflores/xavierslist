package routes

import (
	"github.com/gofiber/fiber/v2"
)

func CategoriesRoutes(a *fiber.App) {
	route := a.Group("/api/v1/categories")

	route.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Get all categories")
	})

	route.Get("/:id", func(c *fiber.Ctx) error {
		return c.SendString("Get category by id")
	})

	route.Get("/:id/subcategories", func(c *fiber.Ctx) error {
		return c.SendString("Get subcategories by category id")
	})


}