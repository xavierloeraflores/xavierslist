package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetAllCategories(c *fiber.Ctx) error {
	return c.SendString("Get all categories")
}

func GetCategoryByCategoryId(c *fiber.Ctx) error {
	categoryId := c.Params("categoryId")
	return c.SendString("Get category by id: " + categoryId)
}

func GetSubcategoriesByCategoryId(c *fiber.Ctx) error {
	categoryId := c.Params("categoryId")
	return c.SendString("Get subcategories by category id: " + categoryId)
}