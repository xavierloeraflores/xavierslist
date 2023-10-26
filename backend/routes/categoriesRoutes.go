package routes

import (
	"github.com/gofiber/fiber/v2"
	"backend/handlers"
)

func CategoriesRoutes(a *fiber.App) {
	route := a.Group("/api/v1/categories")

	route.Get("/", handlers.GetAllCategories)

	route.Get("/:categoryId", handlers.GetCategoryByCategoryId)

	route.Get("/:id/subcategories", handlers.GetSubcategoriesByCategoryId)


}