package controller

import (
	"contact-management-ai/config"
	"contact-management-ai/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CreateAddressRequest struct {
	Street     string `json:"street"`
	City       string `json:"city"`
	Province   string `json:"province"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code"`
}

func CreateAddress(c *fiber.Ctx) error {
	username := c.Locals("username")
	if username == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"errors": "Unauthorized",
		})
	}

	// Ambil contactId dari URL
	contactIdParam := c.Params("contactId")
	contactID, err := strconv.Atoi(contactIdParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": "Invalid contact ID",
		})
	}

	// Pastikan contact milik user
	var contact model.Contact
	err = config.DB.Where("id = ? AND username = ?", contactID, username.(string)).
		First(&contact).Error

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"errors": "Contact is not found",
		})
	}

	var req CreateAddressRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": "Invalid request body",
		})
	}

	// Validasi country wajib
	if req.Country == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": "Country is required",
		})
	}

	address := model.Address{
		ContactID:  contact.ID,
		Street:     req.Street,
		City:       req.City,
		Province:   req.Province,
		Country:    req.Country,
		PostalCode: req.PostalCode,
	}

	if err := config.DB.Create(&address).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": fiber.Map{
			"id":          address.ID,
			"street":      address.Street,
			"city":        address.City,
			"province":    address.Province,
			"country":     address.Country,
			"postal_code": address.PostalCode,
		},
	})
}
