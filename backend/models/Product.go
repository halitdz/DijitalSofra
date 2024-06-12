package models

import (
	"log"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name     string
	Price    float64
	Category string
	Image    string
}

func AddProducts(db *gorm.DB) {
	products := []Product{
		{Name: "Kuşkonmaz", Price: 12.99, Category: "Salad", Image: "/food/food_1.png"},
		{Name: "Ton Balıklı Salata", Price: 15.99, Category: "Pizza", Image: "/food/food_2.png"},
		{Name: "Marul salatası", Price: 18.99, Category: "Pizza", Image: "/food/food_3.png"},
		{Name: "Pepperoni Pizza", Price: 16.99, Category: "Pizza", Image: "/food/food_4.png"},
		{Name: "Caesar Salad", Price: 10.99, Category: "Salad", Image: "/food/food_5.png"},
		{Name: "Beef Burger", Price: 11.99, Category: "Burger", Image: "/food/food_6.png"},
		{Name: "Veggie Burger", Price: 9.99, Category: "Burger", Image: "/food/food_7.png"},
		{Name: "Cheeseburger", Price: 12.49, Category: "Burger", Image: "/food/food_8.png"},
		{Name: "Chicken Wrap", Price: 8.99, Category: "Wrap", Image: "/food/food_9.png"},
		{Name: "Veggie Wrap", Price: 7.99, Category: "Wrap", Image: "/food/food_10.png"},
		{Name: "Sushi Roll", Price: 14.99, Category: "Sushi", Image: "/food/food_11.png"},
		{Name: "Tempura Shrimp", Price: 13.99, Category: "Sushi", Image: "/food/food_12.png"},
		{Name: "California Roll", Price: 12.99, Category: "Sushi", Image: "/food/food_13.png"},
		{Name: "Dragon Roll", Price: 15.99, Category: "Sushi", Image: "/food/food_14.png"},
		{Name: "Beef Tacos", Price: 10.99, Category: "Tacos", Image: "/food/food_15.png"},
		{Name: "Chicken Tacos", Price: 9.99, Category: "Tacos", Image: "/food/food_16.png"},
		{Name: "Fish Tacos", Price: 11.99, Category: "Tacos", Image: "/food/food_17.png"},
		{Name: "Pad Thai", Price: 14.99, Category: "Noodles", Image: "/food/food_18.png"},
		{Name: "Ramen", Price: 13.99, Category: "Noodles", Image: "/food/food_19.png"},
		{Name: "Spaghetti Carbonara", Price: 12.99, Category: "Pasta", Image: "/food/food_20.png"},
		{Name: "Fettuccine Alfredo", Price: 13.49, Category: "Pasta", Image: "/food/food_21.png"},
		{Name: "Chicken Parmesan", Price: 15.99, Category: "Pasta", Image: "/food/food_22.png"},
		{Name: "Lobster Bisque", Price: 17.99, Category: "Soup", Image: "/food/food_23.png"},
		{Name: "Tom Yum Soup", Price: 14.99, Category: "Soup", Image: "/food/food_24.png"},
		{Name: "Minestrone Soup", Price: 9.99, Category: "Soup", Image: "/food/food_25.png"},
		{Name: "Grilled Salmon", Price: 19.99, Category: "Fish", Image: "/food/food_26.png"},
		{Name: "Fish and Chips", Price: 14.99, Category: "Fish", Image: "/food/food_27.png"},
		{Name: "Roasted Chicken", Price: 15.99, Category: "Chicken", Image: "/food/food_28.png"},
		{Name: "Chicken Wings", Price: 12.99, Category: "Chicken", Image: "/food/food_29.png"},
		{Name: "BBQ Ribs", Price: 18.99, Category: "Meat", Image: "/food/food_30.png"},
		{Name: "Steak", Price: 24.99, Category: "Meat", Image: "/food/food_31.png"},
		{Name: "Lamb Chops", Price: 22.99, Category: "Meat", Image: "/food/food_32.png"},
	}

	for _, product := range products {
		var existingProduct Product
		if err := db.Where("name = ? AND price = ? AND category = ? AND image = ?", product.Name, product.Price, product.Category, product.Image).First(&existingProduct).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				db.Create(&product)
			} else {
				log.Printf("Error checking product %v: %v\n", product.Name, err)
			}
		}
	}
}
