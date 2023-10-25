package routes

import (
	"github.com/gofiber/fiber/v2"
	"backend/handlers"
)

func CategoriesRoutes(a *fiber.App) {
	route := a.Group("/api/v1/categories")

	route.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Get Users")
	})

	route.Post("/",func(c *fiber.Ctx) error {
		return c.SendString("Post User")
	})

	route.Get("/user/:userId",func(c *fiber.Ctx) error {
		return c.SendString("Get User by userId")
	})

	route.Put("/user/:userId",func(c *fiber.Ctx) error {
		return c.SendString("Edit User by userId")
	})

	route.Delete("/user/:userId",func(c *fiber.Ctx) error {
		return c.SendString("Delete User by userId")
	})

}