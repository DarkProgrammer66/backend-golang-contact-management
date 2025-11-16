package controller

import (
	"contact-management-ai/config"
	"contact-management-ai/model"
	"net/mail"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UpdateContactRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

func UpdateContact(c *fiber.Ctx) error {
	username := c.Locals("username")
	if username == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"errors": "Unauthorized",
		})
	}

	// Get contact ID
	idParam := c.Params("id")
	contactID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": "Invalid contact ID",
		})
	}

	// Parse request body
	var req UpdateContactRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": "Invalid request body",
		})
	}

	// Validate email
	if req.Email != "" {
		if _, err := mail.ParseAddress(req.Email); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"errors": "Email is not valid format",
			})
		}
	}

	// Find existing contact
	var contact model.Contact
	err = config.DB.
		Where("id = ? AND username = ?", contactID, username.(string)).
		First(&contact).Error

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"errors": "Contact is not found",
		})
	}

	// Update data
	contact.FirstName = req.FirstName
	contact.LastName = req.LastName
	contact.Email = req.Email
	contact.Phone = req.Phone

	// Save to DB
	if err := config.DB.Save(&contact).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	// Success response
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
