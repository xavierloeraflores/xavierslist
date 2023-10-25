package routes

import (
	"github.com/gofiber/fiber/v2"
	"backend/handlers"
)

func PostsRoutes(a *fiber.App) {
	route := a.Group("/api/v1/posts")

	route.Get("/", handlers.GetPosts)

	route.Post("/", handlers.PostPost)

	route.Get("/:postId", handlers.GetPostByPostId)

	route.Put("/:postId", handlers.UpdatePostByPostId)
	
	route.Delete("/:postId", handlers.DeletePostsByPostId)

	route.Get("/category/:categoryId", handlers.GetPostsByCategoryId)

	route.Get("/subcategory/:subcategoryId", handlers.GetPostsBySubcategoryId)

	route.Get("/user/:userId", handlers.GetPostsByUserId)

}