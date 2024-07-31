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
	var selectedProducts []model.Products
	var category model.Categories

	for {
		//display product
		products := showProduct(db)

		if len(products) == 0 {
			fmt.Println("No Product Available.")
			return
		}

		//Prompt user to choose a Product Id
		var productId int
		fmt.Print("Choose Your Product ID: ")
		fmt.Scan(&productId)

		//find select product
		var selectedProduct *model.Products
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

	}
}

// showProduct show list product
func showProduct(db *sql.DB) []model.Products {
	//initialize struct DB
	dbInit := database2.Database{DB: db}
	rows, err := dbInit.ShowProducts()
	var products []model.Products

	if err != nil {
		log.Fatalf("failed to show products: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var product model.Products
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
