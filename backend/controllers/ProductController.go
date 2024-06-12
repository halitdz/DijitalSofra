package controllers

import (
	"go-backend/models"

	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	var products []models.Product
	models.DB.Find(&products)
	return c.JSON(products)
}

func AddProduct(c *fiber.Ctx) error {
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "JSON parse edilemiyor!"})
	}
	if err := models.DB.Create(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Ürün oluşturulamadı."})
	}
	return c.JSON(product)
}
