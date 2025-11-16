package controller

import (
	"contact-management-ai/config"
	"contact-management-ai/model"

	"github.com/gofiber/fiber/v2"
)

func GetAddress(c *fiber.Ctx) error {
	// Parse params
	contactId := c.Params("contactId")
	addressId := c.Params("addressId")

	var contact model.Contact
	if err := config.DB.First(&contact, "id = ?", contactId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"errors": "contact is not found",
		})
	}

	var address model.Address
	if err := config.DB.Where("id = ? AND contact_id = ?", addressId, contactId).
		First(&address).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"errors": "address is not found",
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
