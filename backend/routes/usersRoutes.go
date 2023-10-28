package routes

import (
	"backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func UsersRoutes(a *fiber.App) {
	route := a.Group("/api/v1/users")

	route.Get("/", handlers.GetAllUsers)

	route.Post("/", handlers.PostUser)

	route.Get("/:userId", handlers.GetUserByUserId)

	route.Put("/:userId", handlers.UpdateUserByUserId)

	route.Delete("/:userId", handlers.DeleteUserByUserId)

}