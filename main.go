package main

import (
	"contact-management-ai/config"
	"contact-management-ai/controller"
	"contact-management-ai/handler"
	"contact-management-ai/middleware"
	"contact-management-ai/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	// ✅ Izinkan semua origin (selama pengembangan)
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173", // bisa diganti dengan "http://localhost:5173"
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	config.ConnectDatabase()
	config.DB.AutoMigrate(&model.User{})

	app.Post("/api/users", handler.Register)
	app.Post("/api/users/login", handler.Login)
	app.Get("/api/users/current", middleware.JWTProtected, controller.GetCurrentUser)
	app.Patch("/api/users/current", middleware.JWTProtected, controller.UpdateCurrentUser)
	app.Delete("/api/users/logout", controller.LogoutUser)

	app.Post("/api/contacts", middleware.JWTProtected, controller.CreateContact)
	app.Get("/api/contacts", middleware.JWTProtected, controller.SearchContacts)
	app.Delete("/api/contacts/:id", middleware.JWTProtected, controller.DeleteContact)
	app.Get("/api/contacts/:id", middleware.JWTProtected, controller.GetContact)
	app.Put("/api/contacts/:id", middleware.JWTProtected, controller.UpdateContact)
	app.Post("/api/contacts/:contactId/addresses", middleware.JWTProtected, controller.CreateAddress)
	app.Get("/api/contacts/:contactId/addresses/:addressId", middleware.JWTProtected, controller.GetAddress)
	app.Get("/api/contacts/:contactId/addresses", middleware.JWTProtected, controller.ListAddresses)

	app.Listen(":3000")

}
