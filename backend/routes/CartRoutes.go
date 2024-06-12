package routes

import (
	"go-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func CheckoutRoutes(app *fiber.App) {
	cart := app.Group("/cart")
	cart.Post("/", controllers.GetCartItems)
	cart.Post("/checkout", controllers.Checkout)
}
