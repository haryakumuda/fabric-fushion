package handler

import (
	"database/sql"
	"fabric-fushion/database"
	"fabric-fushion/model"
	"fmt"
	"strconv"
	"strings"
)

// sales reports.
func OrderReports(db *sql.DB) {
	reports, err := database.OrderReports(db)
	if err != nil {
		fmt.Println("Error fetching sales reports:", err)
		return
	}

	// Print header
	fmt.Println("Sales Reports:")
	fmt.Println("Date       | Total Quantity | Total Amount")

	// Print the sales details
	for _, report := range reports {
		orderDate := report["order_date"].(string)
		totalQuantity := report["total_quantity"].(int)
		totalAmount := report["total_amount"].(float64)

		fmt.Printf("%-10s | %-14d | %.2f\n", orderDate, totalQuantity, totalAmount)
	}
}

// StockReports displays stock reports with formatted output.
func StockReports(db *sql.DB) {
	reports, err := database.StockReports(db)
	if err != nil {
		fmt.Println("Error fetching stock reports:", err)
		return
	}

	fmt.Println("Stock Reports:")
	fmt.Println("Product ID | Product Name                   | Category       | Stock Level")
	fmt.Println(strings.Repeat("-", 70))

	for _, report := range reports {
		productID := report["product_id"].(int)
		productName := report["product_name"].(string)
		categoryName := report["category"].(string)
		stock := report["stock"].(int)

		fmt.Printf("%-10d | %-30s | %-14s | %-11d\n", productID, productName, categoryName, stock)
	}
}

func UserReports(db *sql.DB) {
	customers, err := database.ShowCustomer(db)
	if err != nil {
		fmt.Println("Error showing customer: ", err)
		return
	}

	customerMap := make(map[int64]model.ShowCustomer)
	customersId := []int64{0}

	for _, customer := range customers {
		fmt.Printf("%d: %s (%s)\n", customer.CustomerId, customer.Name, customer.Email)
		customerID := int64(customer.CustomerId)
		customersId = append(customersId, customerID)
		customerMap[customerID] = customer
	}

	fmt.Print("Enter customer ID: ")
	var input string
	fmt.Scanln(&input)

	customerID, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid customer ID.")
		return
	}

	reports, err := database.UserReports(db, customerID)
	if err != nil {
		fmt.Println("Error fetching user reports:", err)
		return
	}

	if len(reports) == 0 {
		fmt.Println("No reports available for this customer ID.")
		return
	}

	fmt.Println("User Reports:")
	fmt.Println("Product ID | Product Name                   | Quantity | Total Price | Date       ")
	fmt.Println(strings.Repeat("-", 70))

	for _, report := range reports {
		productID := report["product_id"].(int)
		productName := report["product_name"].(string)
		quantity := report["quantity"].(int)
		totalPrice := report["total_price"].(float64)
		orderDate := report["order_date"].(string)

		fmt.Printf("%-10d | %-30s | %-8d | %.2f       | %-10s\n", productID, productName, quantity, totalPrice, orderDate)
	}
}
