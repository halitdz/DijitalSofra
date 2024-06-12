package controllers

import (
	"go-backend/models"

	"github.com/gofiber/fiber/v2"
)

type OrderItem struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

type OrderRequest struct {
	UserID uint        `json:"user_id"`
	Items  []OrderItem `json:"items"`
}

func PlaceOrder(c *fiber.Ctx) error {
	var orderRequest OrderRequest
	if err := c.BodyParser(&orderRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "JSON parse edilemiyor."})
	}

	for _, item := range orderRequest.Items {
		order := models.Order{
			UserID:    orderRequest.UserID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		}
		if err := models.DB.Create(&order).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Sipariş verilemedi!"})
		}
	}

	return c.JSON(fiber.Map{"message": "Sipariş başarıyla verildi!"})
}
