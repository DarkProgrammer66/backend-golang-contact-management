package controller

import (
	"contact-management-ai/config"
	"contact-management-ai/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func DeleteContact(c *fiber.Ctx) error {
	username := c.Locals("username")
	if username == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"errors": "Unauthorized",
		})
	}

	// Get contact ID from URL params
	idParam := c.Params("id")
	contactID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": "Invalid contact ID",
		})
	}

	var contact model.Contact

	// Ensure contact belongs to the authenticated user
	err = config.DB.Where("id = ? AND username = ?", contactID, username.(string)).First(&contact).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"errors": "Contact is not found",
		})
	}

	// Delete the contact
	if err := config.DB.Delete(&contact).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": "OK",
	})
}
