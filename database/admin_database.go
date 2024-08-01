package database

import (
	"database/sql"
	"fabric-fushion/model"
	"fmt"
	"log"
)

func AddProduct(db *sql.DB, product model.Product) (int64, error) {
	query := "INSERT INTO products (category_id, name, price, stock) VALUES (?, ?, ?, ?)"
	result, err := db.Exec(query, product.CategoryID, product.Name, product.Price, product.Stock)
	if err != nil {
		fmt.Println("Error adding product:", err)
		return 0, err
	}

	// Get the ID of the newly inserted row
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	log.Println("Product added successfully with ID:", id)
	return id, nil
}

func DeleteProduct(db *sql.DB, id int64) (int64, error) {
	query := "DELETE FROM products WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		fmt.Println("Error adding product:", err)
		return 0, err
	}

	log.Println("Product added successfully with ID:", id)
	return id, nil
}

func ShowCategories(db *sql.DB) ([]model.Categories, error) {
	var categories []model.Categories

	// Execute the query
	rows, err := db.Query("SELECT id, category FROM categories")
	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	// Iterate through the results
	for rows.Next() {
		var id int64
		var category string
		if err := rows.Scan(&id, &category); err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		categories = append(categories, model.Categories{ID: id, Category: category})
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		fmt.Println("Error during iteration:", err)
	}

	return categories, nil
}
