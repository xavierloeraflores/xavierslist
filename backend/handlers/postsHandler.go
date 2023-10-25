package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetPosts(c *fiber.Ctx) error {
	return c.SendString("Hello World from posts")
}

func PostPost(c *fiber.Ctx) error {
	return c.SendString("Posted to posts:\n" + string(c.Body()))
}

func GetPostByPostId(c *fiber.Ctx) error {
	postId := c.Params("postId")
	return c.SendString("Get post by id  " + postId)
}

func UpdatePostByPostId(c *fiber.Ctx) error {
	postId := c.Params("postId")
	return c.SendString("Update post by id:\n"+ postId + string(c.Body()))
}

func DeletePostsByPostId(c *fiber.Ctx) error {
	postId := c.Params("postId")
	return c.SendString("Delete post by id: "+ postId)
}

 func GetPostsByCategoryId(c *fiber.Ctx) error {
	categoryId := c.Params("categoryId")
	return c.SendString("Get posts by category id: "+ categoryId)
}

func GetPostsBySubcategoryId(c *fiber.Ctx) error {
	subcategoryId := c.Params("subcategoryId")
	return c.SendString("Get posts by subcategory id: "+ subcategoryId)
}

func GetPostsByUserId(c *fiber.Ctx) error {
	userId := c.Params("userId")
	return c.SendString("Get posts by user id: " + userId)
}