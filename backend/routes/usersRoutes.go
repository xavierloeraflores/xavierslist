package routes

import (
	"backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func UsersRoutes(a *fiber.App) {
	route := a.Group("/api/v1/users")

	route.Get("/", handlers.GetAllUsers)

	route.Post("/", handlers.PostUser)

	route.Get("/user/:userId", handlers.GetUserByUserId)

	route.Put("/user/:userId", handlers.UpdateUserByUserId)

	route.Delete("/user/:userId", handlers.DeleteUserByUserId)

}