package main

import (
	// "log"
	"os"

	"backend/database"
	"backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
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

func loadENV() {
    err := godotenv.Load(".env")
    if err != nil {
        panic("Error loading .env file")
    }
}

func main() {
    loadENV()
    app := fiber.New()

    database.ConnectDB()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })
    routes.PostsRoutes(app)
    routes.CategoriesRoutes(app)
	routes.UsersRoutes(app)

    app.Listen(getPort())
    // log.Fatal(app.Listen(":3000"))
}
