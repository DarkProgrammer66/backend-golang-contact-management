package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte("RAHASIA")

func JWTProtected(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	fmt.Println("Authorization header:", authHeader)

	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"errors": "b",
		})
	}
	tokenStr := strings.Trim(authHeader, "\" ")
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	fmt.Println("Token setelah trim:", tokenStr)

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"errors": "c",
		})
	}

	claims := token.Claims.(jwt.MapClaims)
	c.Locals("username", claims["username"])
	return c.Next()
}
