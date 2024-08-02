package handler

import (
	"database/sql"
	database2 "fabric-fushion/database"
	"fabric-fushion/model"
	"fmt"
	"log"
	"strings"
)

func BuyProduct(db *sql.DB, customerId int) error {
	// Initialize struct db
	dbInit := database2.Database{DB: db}
	var selectedProducts []model.Product
	productQuantities := make(map[int]int)

	for {
		// Display product
		products := ShowProduct(db)

		if len(products) == 0 {
			fmt.Println("No Product Available.")
			return nil
		}

		// Prompt user to choose a productID
		var productId int64
		fmt.Print("Choose Your Product ID: ")
		fmt.Scan(&productId)

		// Find selected product
		var selectedProduct *model.Product
		for _, product := range products {
			if product.ID == productId {
				selectedProduct = &product
				break
			}
		}
		if selectedProduct == nil {
			fmt.Println("Product Not Found. Please choose a valid Product ID.")
			continue
		}

		// Add selected Product to the list
		selectedProducts = append(selectedProducts, *selectedProduct)
		fmt.Printf("You Selected Id: %d\n Name: %s\n Price: %.2f\n Category: %s\n",
			selectedProduct.ID, selectedProduct.Name, selectedProduct.Price, selectedProduct.Category)

		productQuantities[int(selectedProduct.ID)] = 1

		// Ask if user wants to buy more products
		var choice string
		fmt.Print("Do you want to buy another product? (y/n): ")
		fmt.Scan(&choice)
		choice = strings.ToLower(choice)
		if choice != "y" && choice != "yes" {
			break
		}
	}

	if len(selectedProducts) == 0 {
		fmt.Println("No Product Selected.")
		return nil
	}

	// stok product -1
	for _, product := range selectedProducts {
		if err := dbInit.UpdateProductStock(product.ID, -1); err != nil {
			fmt.Printf("Failed to update stock for Product ID %d: %v\n", product.ID, err)
			return nil
		}
	}

	// Call store procedure to save data sales
	err := dbInit.InsertSale(customerId, productQuantities)
	if err != nil {
		fmt.Printf("Failed to insert into table sales: %v\n", err)
		return nil
	}

	// Display the products that have been purchased
	fmt.Println("\n=============Products you have purchased:===================")
	for _, product := range selectedProducts {
		fmt.Printf("ID: %d | Name: %s | Price: %.2f | Category: %s\n",
			product.ID, product.Name, product.Price, product.Category)
	}
	fmt.Println("=============Thank you for your purchase!==============")
	return nil
}

// ShowProduct show list product
func ShowProduct(db *sql.DB) []model.Product {
	//initialize struct DB
	dbInit := database2.Database{DB: db}
	rows, err := dbInit.ShowProducts()
	var products []model.Product

	if err != nil {
		log.Fatalf("failed to show products: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var product model.Product
		var category string
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &category); err != nil {
			log.Fatalf("failed to scan product : %v", err)
		}
		product.Category = category
		fmt.Printf("ID : %d | Name : %s | Price : %.2f | Category : %s\n", product.ID, product.Name, product.Price, product.Category)
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return products
}

// OrderHistory show orderHistory Data
func OrderHistory(db *sql.DB, customerID int) {
	//initialize Database
	//format Date
	dateFormat := "2006-01-02 15:04:05"
	dbInit := database2.Database{DB: db}

	orders, err := dbInit.GetOrderHistory(customerID)
	if err != nil {
		log.Fatalf("Error querying orders for customerID %d: %v", customerID, err)
	}

	for _, order := range orders {
		productRows, err := dbInit.GetProductsForOrder(order.ID)
		if err != nil {
			log.Fatalf("Error getting products for orderID %d: %v", order.ID, err)
		}
		defer productRows.Close()

		//formatDate
		formattedDate := order.OrderDate.Format(dateFormat)
		fmt.Printf("Order ID: %d, Order Date: %s\n", order.ID, formattedDate)
		for productRows.Next() {
			var detail model.SalesProductDetail
			if err := productRows.Scan(&detail.ID, &detail.SaleID, &detail.ProductID, &detail.Quantity, &detail.Name, &detail.Price, &detail.CategoryID, &detail.Category); err != nil {
				log.Fatalf("Error scanning product detail: %v", err)
			}
			fmt.Printf("Product ID: %d, Name: %s, Quantity: %d, Price: %.2f, Category: %s\n",
				detail.ProductID, detail.Name, detail.Quantity, detail.Price, detail.Category)
		}

		if err := productRows.Err(); err != nil {
			log.Fatalf("Error iterating over product rows: %v", err)
		}

		fmt.Println()
	}
}
