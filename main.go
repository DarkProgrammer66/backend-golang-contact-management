package main

import (
	"contact-management-ai/config"
	"contact-management-ai/handler"
	"contact-management-ai/model"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.ConnectDatabase()
	config.DB.AutoMigrate(&model.User{})

	app.Post("/api/users", handler.Register)

	app.Listen(":3000")
}
