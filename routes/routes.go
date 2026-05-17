package routes

import (
	"go-auth/handlers"
	"go-auth/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")

	// Public routes
	auth := api.Group("/auth")
	api.Post("/signup", handlers.Signup)
	api.Post("/login", handlers.Login)

	// Protected routes
	user := api.Group("/user", middleware.Protected)
	user.Get("/me", handlers.Me)
}