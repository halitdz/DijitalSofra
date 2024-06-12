package main

import (
	"fmt"
	"go-backend/models"
	"go-backend/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDatabase() {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=mydb port=4321 sslmode=disable "
	models.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!")
		fmt.Println(err)
	}

	log.Println("Database connection successfully opened.")

	models.DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})
	models.AddProducts(models.DB)
}

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3001", // Vite default port
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	initDatabase()

	routes.AuthRoutes(app)
	routes.ProductRoutes(app)
	routes.OrderRoutes(app)
	routes.CheckoutRoutes(app)
	routes.UserRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
