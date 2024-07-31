package handler

import (
	"database/sql"
	database2 "fabric-fushion/database"
	"fabric-fushion/model"
	"fmt"
	"log"
)

func BuyProduct(db *sql.DB, customerID int) {
	//initialize struct db
	dbInit := database2.Database{DB: db}
	var selectedProducts []model.Product
	var category model.Categories

	for {
		//display product
		products := showProduct(db)

		if len(products) == 0 {
			fmt.Println("No Product Available.")
			return
		}

		//Prompt user to choose a productID
		var productId uint
		fmt.Print("Choose Your Product ID: ")
		fmt.Scan(&productId)

		//find select product
		var selectedProduct *model.Product
		for _, product := range products {
			if product.ID == productId {
				selectedProduct = &product
				break
			}
		}
		if selectedProduct == nil {
			fmt.Println("Product Not Found . Please choose a valid Product ID . ")
			continue
		}
		// Add selected Product to the list
		selectedProducts = append(selectedProducts, *selectedProduct)
		fmt.Printf("You Selected Id : %d\n Name : %s\n Price : %.2f Category : %s", selectedProduct.ID, selectedProduct.Name, selectedProduct.Price, category.Category)

		// Ask if user wants to buy more product
		var choice string
		fmt.Print("Do you want to buy another product ? (yes/no) : ")
		fmt.Scan(&choice)
		if choice != "yes" {
			break
		}
	}
	if len(selectedProducts) == 0 {
		fmt.Println("No Product Selected . ")
		return
	}

	productQuantities := make(map[int]int)
	for _, product := range selectedProducts {
		productQuantities[int(product.ID)]++
	}

	//call store procedure for save data sales
	err := dbInit.InsertSale(customerID, productQuantities)
	if err != nil {
		fmt.Printf("Failed insert To table sales : %v", err)
		return
	}

	// Display the products that have been purchased
	fmt.Println("\nProducts you have purchased:")
	for _, product := range selectedProducts {
		fmt.Printf("ID: %d\n Name: %s\n Price: %.2f\n Category: %s\n",
			product.ID, product.Name, product.Price, category.Category)
	}
	fmt.Println("Thank you for your purchase!")
}

// showProduct show list product
func showProduct(db *sql.DB) []model.Product {
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
		var category model.Categories
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &category.Category); err != nil {
			log.Fatalf("failed to scan product : %v", err)
		}
		fmt.Printf("ID : %d\n Name : %s\n Price : %.2f\n Category : %s\n", product.ID, product.Name, product.Price, category.Category)
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return products
}
