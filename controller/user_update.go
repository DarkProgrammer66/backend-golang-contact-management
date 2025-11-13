package controller

import (
	"contact-management-ai/config"
	"contact-management-ai/model"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UpdateUserRequest struct {
	Name     *string `json:"name"`
	Password *string `json:"password"`
}

func UpdateCurrentUser(c *fiber.Ctx) error {
	db := config.DB

	// get username from JWT middleware
	usernameIfc := c.Locals("username")
	username, ok := usernameIfc.(string)
	if !ok || username == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"errors": "Unauthorized",
		})
	}

	// find user
	var user model.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"errors": "Unauthorized",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"errors": "Failed to fetch user",
		})
	}

	// parse body
	var req UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": "Invalid request body",
		})
	}

	// validate & update name (optional)
	if req.Name != nil {
		if len(*req.Name) > 100 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"errors": "Name length max 100",
			})
		}
		user.Name = *req.Name
	}

	// update password (optional)
	if req.Password != nil {
		if len(*req.Password) < 6 {
			// optional: enforce minimal length
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"errors": "Password minimum length 6",
			})
		}
		hashed, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"errors": "Failed to hash password",
			})
		}
		user.Password = string(hashed)
	}

	// save
	if err := db.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errors": "Failed to update user",
		})
	}

	// success response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"username": user.Username,
			"name":     user.Name,
		},
	})
}
