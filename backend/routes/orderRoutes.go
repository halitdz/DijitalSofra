package routes

import (
	"go-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func OrderRoutes(app *fiber.App) {
	order := app.Group("/order")

	order.Post("/order", controllers.PlaceOrder)
}
