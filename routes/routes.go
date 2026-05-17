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
	auth.Post("/signup", handlers.Signup)
	auth.Post("/login", handlers.Login)

	// Protected routes
	user := api.Group("/user", middleware.Protected)
	user.Get("/me", handlers.Me)
}