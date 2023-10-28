package handlers

import (
	"backend/database"
	"backend/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	users := []models.User{}
	database.DB.Find(&users)

	if (len(users) == 0) {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"message": "No users found",
			"data": nil,
		})

	}

	return c.JSON(fiber.Map{
		"status": "success",
		"message": "All users found",
		"data": users,
	})
}

func PostUser(c *fiber.Ctx) error {
	requestBody := c.Body()
	return c.SendString("Post User: \n" + string(requestBody))
}

func GetUserByUserId(c *fiber.Ctx) error {
	userId := c.Params("userId")
	user := models.User{}
	result := database.DB.First(&user,"ID = ?", userId)
	
	if (result.Error != nil) {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"message": "User not found",
			"data": nil,
		})
	}

	return c.JSON(fiber.Map{
		"status": "success", 
		"message": "User found", "data": user,
	})
}

func UpdateUserByUserId(c *fiber.Ctx) error {
	userId := c.Params("userId")
	requestBody := c.Body()
	return c.SendString("Edit User by userId: " + userId + string(requestBody))
}

func DeleteUserByUserId(c *fiber.Ctx) error {
	userId := c.Params("userId")
	user := models.User{}
	result := database.DB.First(&user,"ID = ?", userId)
	
	if (result.Error != nil) {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"message": "User not found",
			"data": nil,
		})
	}

	err := database.DB.Delete(&user, "id=?", userId).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"message": "User not deleted",
			"data": nil,
		})
	}

	return c.JSON(fiber.Map{
		"status": "success", 
		"message": "User deleted",
	})
}