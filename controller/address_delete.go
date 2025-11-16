package controller

import (
	"contact-management-ai/config"
	"contact-management-ai/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func DeleteAddress(c *fiber.Ctx) error {
	contactIdStr := c.Params("id")
	addressIdStr := c.Params("addressId")

	// Convert to uint
	contactId, err := strconv.ParseUint(contactIdStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": "invalid contact id",
		})
	}

	addressId, err := strconv.ParseUint(addressIdStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": "invalid address id",
		})
	}

	var contact model.Contact
	if err := config.DB.First(&contact, contactId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"errors": "contact is not found",
		})
	}

	var address model.Address
	if err := config.DB.
		Where("id = ? AND contact_id = ?", addressId, contactId).
		First(&address).Error; err != nil {

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"errors": "address is not found",
		})
	}

	if err := config.DB.Delete(&address).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data": "OK",
	})
}
