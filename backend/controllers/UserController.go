package controllers

import (
	"errors"
	"go-backend/models"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashedPassword)
	if err := models.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create user"})
	}
	return c.JSON(user)
}

func LoginUser(c *fiber.Ctx) error {
	input := new(LoginInput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "JSON parse edilemiyor."})
	}
	var user models.User
	if err := models.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz e-posta veya şifre!"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Kullanıcı girişi yapılamadı!"})
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz e-posta veya şifre!"})
	}
	token, err := createToken(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Token oluşturulamadı!"})
	}
	return c.JSON(fiber.Map{"token": token})
}

func createToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func GetUser(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	var user models.User
	if err := models.DB.First(&user, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Kullanıcı bulunamadı!"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database hatası."})
	}
	response := fiber.Map{
		"id":      user.ID,
		"name":    user.Name,
		"email":   user.Email,
		"balance": user.Balance,
	}

	return c.JSON(response)
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

func GenerateJWT(userID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func getUserIDFromToken(c *fiber.Ctx) (uint, error) {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return 0, errors.New("missing authorization header")
	}

	tokenString := authHeader
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return 0, errors.New("Geçersiz token")
	}

	return claims.UserID, nil
}
