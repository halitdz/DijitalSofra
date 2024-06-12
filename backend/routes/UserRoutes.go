package routes

import (
	"go-backend/controllers"
	middlewares "go-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	user := app.Group("/user")
	user.Use(middlewares.JWTMiddleware())
	user.Get("/", controllers.GetUser)

}
