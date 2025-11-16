package controller

import (
	"contact-management-ai/config"
	"contact-management-ai/model"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func SearchContacts(c *fiber.Ctx) error {
	username := c.Locals("username")
	if username == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"errors": "Unauthorized",
		})
	}

	name := c.Query("name")
	email := c.Query("email")
	phone := c.Query("phone")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	size, _ := strconv.Atoi(c.Query("size", "10"))

	var contacts []model.Contact

	// FIX PENTING: FILTER BERDASARKAN USERNAME
	query := config.DB.Where("username = ?", username.(string))

	if name != "" {
		like := "%" + name + "%"
		query = query.Where("first_name LIKE ? OR last_name LIKE ?", like, like)
	}
	if email != "" {
		query = query.Where("email LIKE ?", "%"+email+"%")
	}
	if phone != "" {
		query = query.Where("phone LIKE ?", "%"+phone+"%")
	}

	var total int64
	query.Model(&model.Contact{}).Count(&total)

	offset := (page - 1) * size

	query.Limit(size).Offset(offset).Find(&contacts)

	return c.JSON(fiber.Map{
		"data": contacts,
		"paging": fiber.Map{
			"page":       page,
			"total_page": int(math.Ceil(float64(total) / float64(size))),
			"total_item": total,
		},
	})
}
