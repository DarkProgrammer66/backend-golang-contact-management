package controller

import (
	"contact-management-ai/config"
	"contact-management-ai/model"

	"github.com/gofiber/fiber/v2"
)

func GetCurrentUser(c *fiber.Ctx) error {
	username := c.Locals("username").(string)

	var user model.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"errors": "a",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"username": user.Username,
			"name":     user.Name,
		},
	})
}
