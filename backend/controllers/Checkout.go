package controllers

import (
	"go-backend/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type CheckoutRequest struct {
	CartItems map[string]int `json:"cartItems"`
}

func Checkout(c *fiber.Ctx) error {

	userID := c.Locals("userID")

	var user models.User
	if err := models.DB.First(&user, userID).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Kullanıcı Bulunamadı."})
	}

	var req CheckoutRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "JSON parse edilemiyor"})
	}

	var totalAmount float64
	for productID, quantity := range req.CartItems {
		var product models.Product
		if err := models.DB.First(&product, productID).Error; err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Product bulunamadı."})
		}
		if quantity <= 0 {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz miktar!"})
		}

		totalAmount += product.Price * float64(quantity)
	}
	totalAmount = totalAmount + 2

	if user.Balance < totalAmount {
		return c.Status(http.StatusPaymentRequired).JSON(fiber.Map{"error": "Yetersiz bakiye!"})
	}

	// Kullanıcı bakiyesini düşür
	user.Balance -= totalAmount
	if err := models.DB.Save(&user).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Kullanıcı bakiyesi güncellenemedi."})
	}

	return c.JSON(fiber.Map{"message": "Checkout successful"})
}
