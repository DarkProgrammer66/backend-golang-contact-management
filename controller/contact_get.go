package controller

import (
	"contact-management-ai/config"
	"contact-management-ai/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetContact(c *fiber.Ctx) error {
	username := c.Locals("username")
	if username == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"errors": "Unauthorized",
		})
	}

	// Ambil ID dari URL
	idParam := c.Params("id")
	contactID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": "Invalid contact ID",
		})
	}

	var contact model.Contact

	// Cari contact yang ID-nya sesuai dan milik user
	err = config.DB.Where("id = ? AND username = ?", contactID, username.(string)).
		First(&contact).Error

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"errors": "Contact is not found",
		})
	}

	return c.JSON(fiber.Map{
		"data": fiber.Map{
			"id":         contact.ID,
			"first_name": contact.FirstName,
			"last_name":  contact.LastName,
			"email":      contact.Email,
			"phone":      contact.Phone,
		},
	})
}
