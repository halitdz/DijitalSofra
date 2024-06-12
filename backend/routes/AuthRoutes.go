package routes

import (
	"go-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	auth := app.Group("/auth")

	auth.Post("/register", controllers.RegisterUser)
	auth.Post("/login", controllers.LoginUser)

}
