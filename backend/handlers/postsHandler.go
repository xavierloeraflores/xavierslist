package handlers

import (
	"backend/database"
	"backend/models"
	"fmt"

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
	post := models.Post{}
	result := database.DB.First(&post,"ID = ?", postId)
	
	if (result.Error != nil) {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"message": "Post not found",
			"data": nil,
		})
	}

	return c.JSON(fiber.Map{
		"status": "success", 
		"message": "User found", "data": post,
	})

}

func UpdatePostByPostId(c *fiber.Ctx) error {
	postId := c.Params("postId")
	return c.SendString("Update post by id:\n"+ postId + string(c.Body()))
}

func DeletePostsByPostId(c *fiber.Ctx) error {
	postId := c.Params("postId")
	post := models.Post{}
	result := database.DB.First(&post,"ID = ?", postId)
	
	if (result.Error != nil) {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"message": "Post not found",
			"data": nil,
		})
	}

	err := database.DB.Delete(&post).Error

	if (err != nil) {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"message": "Post not deleted",
		})
	}

	return c.JSON(fiber.Map{
		"status": "success", 
		"message": "Post deleted",
	})
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

	user := models.User{}
	result := database.DB.First(&user,"ID = ?", userId)
	
	if (result.Error != nil) {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"message": "User not found",
			"data": nil,
		})
	}

	posts := user.Posts
	result = database.DB.Find(&posts)

	if (result.Error != nil){
		fmt.Print(result.Error)
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"message": "Could not retrieve posts",
			"data": nil,
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"message": "Posts found",
		"data": posts,
	})
}