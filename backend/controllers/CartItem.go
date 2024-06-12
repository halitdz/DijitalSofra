package controllers

import (
	"go-backend/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetCartItems(c *fiber.Ctx) error {
	var cartData struct {
		CartItems map[string]int `json:"cartItems"`
	}
	if err := c.BodyParser(&cartData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	var response []map[string]interface{}
	for productIDStr, quantity := range cartData.CartItems {
		productID, err := strconv.Atoi(productIDStr)
		if err != nil {
			continue
		}

		var product models.Product
		if err := models.DB.First(&product, productID).Error; err != nil {
			if err != gorm.ErrRecordNotFound {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not get product details"})
			}
			continue
		}

		response = append(response, map[string]interface{}{
			"product_id": product.ID,
			"name":       product.Name,
			"price":      product.Price,
			"quantity":   quantity,
			"image":      product.Image, // Eğer ürün resmi varsa
		})
	}

	return c.JSON(response)
}
