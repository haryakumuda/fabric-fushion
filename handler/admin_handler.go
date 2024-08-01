package handler

import (
	"database/sql"
	"fabric-fushion/database"
	"fabric-fushion/model"
	"fmt"
	"strconv"
)

func AddProduct(db *sql.DB) {
	var product model.Product

	categories, err := database.ShowCategories(db)
	if err != nil {
		fmt.Println("Error showing categories: ", err)
		return
	}

	fmt.Println("Enter Product Name:")
	fmt.Scanln(&product.Name)

	var categoriesId []int64

	for _, category := range categories {
		fmt.Printf("%d: %s\n", category.ID, category.Category)
		categoriesId = append(categoriesId, int64(category.ID))
	}
	fmt.Println("Enter Product Category ID:")
	for {
		var inputID int64
		fmt.Scanln(&inputID)

		if contains(categoriesId, inputID) {
			product.CategoryID = inputID
			break
		} else {
			fmt.Println("Please Input correct ID")
		}
	}
	fmt.Println("Enter Product Price:")
	for {
		var input string
		fmt.Scanln(&input)

		// Try to convert the input to float64
		price, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Please input a valid amount")
		} else {
			product.Price = price
			break
		}
	}

	fmt.Println("Enter Product Stock Quantity:")

	for {
		var input string
		fmt.Scanln(&input)

		// Try to convert the input to int64
		stock, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			fmt.Println("Please input a valid amount")
		} else {
			product.Stock = stock
			break
		}
	}

	database.AddProduct(db, product)

	fmt.Println("Product added successfully!")
}

// Helper function to check if an ID exists in a slice
func contains(slice []int64, id int64) bool {
	for _, v := range slice {
		if v == id {
			return true
		}
	}
	return false
}
