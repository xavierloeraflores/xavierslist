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

	route.Get("/:postId", func(c *fiber.Ctx) error {
		postId := c.Params("postId")
		return c.SendString("Get post by id  " + postId)
	})

	route.Put("/:postId", func(c *fiber.Ctx) error {
		postId := c.Params("postId")
		return c.SendString("Update post by id:\n"+ postId + string(c.Body()))
	})
	
	route.Delete("/:postId", func(c *fiber.Ctx) error {
		postId := c.Params("postId")
		return c.SendString("Delete post by id: "+ postId + c.Params("id"))
	})

	route.Get("/category/:categoryId", func(c *fiber.Ctx) error {
		categoryId := c.Params("categoryId")
		return c.SendString("Get posts by category id: "+ categoryId + c.Params("categoryId"))
	})

	route.Get("/subcategory/:subcategoryId", func(c *fiber.Ctx) error {
		subcategoryId := c.Params("subcategoryId")
		return c.SendString("Get posts by subcategory id: "+ subcategoryId + c.Params("subcategoryId"))
	})

	route.Get("/user/:userId", func(c *fiber.Ctx) error {
		userId := c.Params("userId")
		return c.SendString("Get posts by user id: " + userId + c.Params("userId"))
	})

}