package handler

import (
	"database/sql"
	"fabric-fushion/model"
	"fmt"
)

func AddProduct(db *sql.DB) {
	var product model.Product

	fmt.Println("Enter Product Name:")
	fmt.Scanln(&product.Name)
	fmt.Println("Enter Product Category ID:")
	fmt.Scanln(&product.CategoryID)
	fmt.Println("Enter Product Price:")
	fmt.Scanln(&product.Price)
	fmt.Println("Enter Product Stock Quantity:")
	fmt.Scanln(&product.Stock)

	query := "INSERT INTO products (category_id, name, price, stock) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(query, product.CategoryID, product.Name, product.Price, product.Stock)
	if err != nil {
		fmt.Println("Error adding product:", err)
		return
	}

	fmt.Println("Product added successfully!")
}
