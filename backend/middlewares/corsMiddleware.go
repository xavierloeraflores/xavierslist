package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func GetCORS() func(*fiber.Ctx) error{
	corsConfig := cors.Config{
		AllowOrigins: "https://xavierslist.vercel.app",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowOriginsFunc: func(origin string) bool {
			return os.Getenv("ENVIRONMENT") == "development" 
			},
	}
	CORS := cors.New(corsConfig)
	return CORS
}
