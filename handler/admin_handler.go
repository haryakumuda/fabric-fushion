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

func OrderReports(db *sql.DB) {
	rows, err := db.Query("SELECT o.id, o.customer_id, c.name, sp.product_id, p.name, sp.quantity, sp.total_price, o.order_date FROM sales o JOIN customers c ON o.customer_id = c.id JOIN sales_products sp ON o.id = sp.sale_id JOIN products p ON sp.product_id = p.id")
	if err != nil {
		fmt.Println("Error fetching orders:", err)
		return
	}
	defer rows.Close()

	fmt.Println("Order Reports:")
	fmt.Println("Order ID\tCustomer ID\tCustomer Name\tProduct ID\tProduct Name\tQuantity\tTotal Price\tOrder Date")
	for rows.Next() {
		var orderID, customerID, productID, quantity int
		var customerName, productName, orderDate string
		var totalPrice float64
		err := rows.Scan(&orderID, &customerID, &customerName, &productID, &productName, &quantity, &totalPrice, &orderDate)
		if err != nil {
			fmt.Println("Error reading order:", err)
			return
		}
		fmt.Printf("%d\t%d\t%s\t%d\t%s\t%d\t%.2f\t%s\n", orderID, customerID, customerName, productID, productName, quantity, totalPrice, orderDate)
	}
}

func StockReports(db *sql.DB) {
	rows, err := db.Query("SELECT id, name, category_id, stock FROM products")
	if err != nil {
		fmt.Println("Error fetching stock:", err)
		return
	}
	defer rows.Close()

	fmt.Println("Stock Reports:")
	fmt.Println("Product ID\tProduct Name\tCategory ID\tStock")
	for rows.Next() {
		var productID, categoryID, stock int
		var productName string
		err := rows.Scan(&productID, &productName, &categoryID, &stock)
		if err != nil {
			fmt.Println("Error reading product:", err)
			return
		}
		fmt.Printf("%d\t%s\t%d\t%d\n", productID, productName, categoryID, stock)
	}
}
