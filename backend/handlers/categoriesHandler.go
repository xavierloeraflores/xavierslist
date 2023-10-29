package handlers

import (
	"backend/database"
	"backend/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllCategories(c *fiber.Ctx) error {
	categories := []models.Category{}
	result := database.DB.Find(&categories)

	if (result.Error != nil) {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"message": "No categories found",
			"data": nil,
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"message": "All categories found",
		"data": categories,
	})
}

func GetCategoryByCategoryId(c *fiber.Ctx) error {
	categoryId := c.Params("categoryId")
	category := models.Category{}
	result := database.DB.First(&category,"ID = ?", categoryId)

	if (result.Error != nil) {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"message": "Category not found",
			"data": nil,
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"message": "Category found",
		"data": category,
	})
}

func GetSubcategoriesByCategoryId(c *fiber.Ctx) error {
	categoryId := c.Params("categoryId")
	category := models.Category{}
	result := database.DB.First(&category,"ID = ?", categoryId)


	if (result.Error != nil) {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"message": "Category not found",
			"data": nil,
		})
	}
	subcategories := category.Subcategories
	result = database.DB.Find(&subcategories)

	return c.JSON(fiber.Map{
		"status": "success",
		"message": "Category found",
		"data": subcategories,
	})
}