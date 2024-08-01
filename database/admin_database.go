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
		fmt.Println("Error deleting product:", err)
		return 0, err
	}

	log.Println("Product deleted successfully")
	return id, nil
}

func ShowProducts(db *sql.DB) ([]model.Product, error) {

	var products []model.Product

	// Execute the query
	rows, err := db.Query("SELECT p.id AS product_id, p.category_id as category_id, p.name AS product_name,  p.price AS product_price, p.stock AS product_stock, c.category AS category_name FROM products p JOIN categories c ON p.category_id = c.id ORDER BY p.id;")
	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	// Iterate through the results
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.CategoryID, &product.Name, &product.Price, &product.Stock, &product.Category); err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		products = append(products, product)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		fmt.Println("Error during iteration:", err)
	}

	return products, nil

}

func ShowCategories(db *sql.DB) ([]model.Categories, error) {
	var categories []model.Categories

	// Execute the query
	rows, err := db.Query("SELECT id, category FROM categories ORDER BY id")
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

func ShowCustomer(db *sql.DB) ([]model.ShowCustomer, error) {

	var customers []model.ShowCustomer

	// Execute the query
	rows, err := db.Query("SELECT c.id AS customer_id, c.user_id, c.name, u.email FROM customers c JOIN  users u ON c.user_id = u.id ORDER BY c.user_id;")
	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	// Iterate through the results
	for rows.Next() {
		var customer model.ShowCustomer
		if err := rows.Scan(&customer.CustomerId, &customer.UserId, &customer.Name, &customer.Email); err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		customers = append(customers, customer)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		fmt.Println("Error during iteration:", err)
	}

	return customers, nil
}

func DeleteCustomer(db *sql.DB, id int64) (int64, error) {
	query := "DELETE FROM customers WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		fmt.Println("Error deleting customer:", err)
		return 0, err
	}

	log.Println("Customer deleted successfully")
	return id, nil
}

func DeleteUser(db *sql.DB, id int64) (int64, error) {
	query := "DELETE FROM users WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		fmt.Println("Error deleting user:", err)
		return 0, err
	}

	log.Println("User deleted successfully")
	return id, nil
}
