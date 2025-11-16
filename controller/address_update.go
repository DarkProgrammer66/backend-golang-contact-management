package controller

import (
	"contact-management-ai/config"
	"contact-management-ai/model"

	"github.com/gofiber/fiber/v2"
)

type UpdateAddressRequest struct {
	Street     string `json:"street"`
	City       string `json:"city"`
	Province   string `json:"province"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code"`
}

func UpdateAddress(c *fiber.Ctx) error {
	contactId := c.Params("contactId")
	addressId := c.Params("addressId")

	// Check contact existence
	var contact model.Contact
	if err := config.DB.First(&contact, "id = ?", contactId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"errors": "contact is not found",
		})
	}

	// Check address existence
	var address model.Address
	if err := config.DB.Where("id = ? AND contact_id = ?", addressId, contactId).First(&address).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"errors": "address is not found",
		})
	}

	// Parse body
	var req UpdateAddressRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": "Invalid request body",
		})
	}

	// Validation
	if req.Country == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": "Country is required",
		})
	}

	// Update fields
	address.Street = req.Street
	address.City = req.City
	address.Province = req.Province
	address.Country = req.Country
	address.PostalCode = req.PostalCode

	if err := config.DB.Save(&address).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	// Response
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
