package handler

import (
	"database/sql"
	"fmt"
)

func OrderReports(db *sql.DB) {
	// Fetch the order details including product names, quantities, and customer names
	rows, err := db.Query(`
        SELECT sales.id AS order_id, 
               customers.name AS customer_name, 
               sales.order_date, 
               products.name AS product_name, 
               sales_products.quantity, 
               (sales_products.quantity * products.price) AS total_amount
        FROM sales
        JOIN customers ON sales.customer_id = customers.id
        JOIN sales_products ON sales.id = sales_products.sale_id
        JOIN products ON sales_products.product_id = products.id
        ORDER BY sales.order_date DESC
    `)
	if err != nil {
		fmt.Println("Error fetching order reports:", err)
		return
	}
	defer rows.Close()

	// Print header
	fmt.Println("Order Reports:")
	fmt.Println("Order ID | Customer Name | Date | Product | Quantity | Total Amount")

	// Iterate through rows
	for rows.Next() {
		var orderID int
		var customerName, orderDate, productName string
		var quantity int
		var totalAmount float64

		// Scan the row into variables
		err := rows.Scan(&orderID, &customerName, &orderDate, &productName, &quantity, &totalAmount)
		if err != nil {
			fmt.Println("Error scanning order report:", err)
			return
		}

		// Print the order details
		fmt.Printf("%d | %s | %s | %s | %d | %.2f\n", orderID, customerName, orderDate, productName, quantity, totalAmount)
	}
}

func StockReports(db *sql.DB) {
	rows, err := db.Query(`
		select products.id, products.name, categories.category, products.stock
		from products
		join categories on products.category_id = categories.id
		order by products.name
	`)
	if err != nil {
		fmt.Println("Error fetching stock reports:", err)
		return
	}
	defer rows.Close()

	fmt.Println("Stock Reports:")
	fmt.Println("Product ID | Product Name | Category | Stock Level")
	for rows.Next() {
		var productID, stock int
		var productName, categoryName string
		err := rows.Scan(&productID, &productName, &categoryName, &stock)
		if err != nil {
			fmt.Println("Error scanning stock report:", err)
			return
		}
		fmt.Printf("%d | %s | %s | %d\n", productID, productName, categoryName, stock)
	}
}
