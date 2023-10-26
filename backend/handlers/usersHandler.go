package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	return c.SendString("Get Users")
}

func PostUser(c *fiber.Ctx) error {
	requestBody := c.Body()
	return c.SendString("Post User: \n" + string(requestBody))
}

func GetUserByUserId(c *fiber.Ctx) error {
	userId := c.Params("userId")
	return c.SendString("Get User by userId: " + userId)
}

func UpdateUserByUserId(c *fiber.Ctx) error {
	userId := c.Params("userId")
	requestBody := c.Body()
	return c.SendString("Edit User by userId: " + userId + string(requestBody))
}

func DeleteUserByUserId(c *fiber.Ctx) error {
	userId := c.Params("userId")
	return c.SendString("Delete User by userId: " + userId)
}