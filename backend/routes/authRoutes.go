package routes

import (
	"backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(a *fiber.App){
	route := a.Group("/api/v1/auth")

	route.Post("/register", handlers.RegisterUser)

	route.Post("/login", handlers.LoginUser)
	

}
