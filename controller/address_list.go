package controller

import (
	"contact-management-ai/config"
	"contact-management-ai/model"

	"github.com/gofiber/fiber/v2"
)

func ListAddresses(c *fiber.Ctx) error {
	contactId := c.Params("id")

	// Cek apakah contact ada
	var contact model.Contact
	if err := config.DB.First(&contact, "id = ?", contactId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"errors": "contact is not found",
		})
	}

	// Ambil semua address milik contact
	var addresses []model.Address
	if err := config.DB.Where("contact_id = ?", contactId).Find(&addresses).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	// Siapkan response list
	var response []fiber.Map
	for _, addr := range addresses {
		response = append(response, fiber.Map{
			"id":          addr.ID,
			"street":      addr.Street,
			"city":        addr.City,
			"province":    addr.Province,
			"country":     addr.Country,
			"postal_code": addr.PostalCode,
		})
	}

	return c.JSON(fiber.Map{
		"data": response,
	})
}
