package main

import (
	// "log"
	"os"

	"backend/routes"

	"github.com/gofiber/fiber/v2"
)



func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	return port
}

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })
    routes.PostsRoutes(app)

    app.Listen(getPort())
    // log.Fatal(app.Listen(":3000"))
}
