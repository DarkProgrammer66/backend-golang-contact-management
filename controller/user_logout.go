package controller

import (
	"contact-management-ai/model"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

var SecretKey = []byte("RAHASIA")

func LogoutUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Ambil token dari header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"errors": "Unauthorized",
			})
		}

		tokenStr := strings.TrimSpace(authHeader)

		// Parse token
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return SecretKey, nil
		})
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"errors": "Unauthorized",
			})
		}

		// Ambil username dari claims
		claims := token.Claims.(jwt.MapClaims)
		username := claims["username"].(string)

		// Cari user berdasarkan username
		var user model.User
		if err := db.Where("username = ?", username).First(&user).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"errors": "Unauthorized",
			})
		}

		// Kosongkan token
		user.Token = ""
		if err := db.Save(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"errors": "Failed to logout",
			})
		}

		// Respon sukses
		return c.JSON(fiber.Map{
			"data": "OK",
		})
	}
}
