package routes

import (
	"go-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(app *fiber.App) {
	product := app.Group("/product")
	product.Get("/products", controllers.GetProducts)
	// product.Post("/product", controllers.AddProduct)
}
