package handler

import (
	"database/sql"
	"fabric-fushion/model"
	"fmt"
	"log"
	"time"
)

func BuyProduct(db *sql.DB, customerID int) {
	products := showProduct(db)

	if len(products) == 0 {
		fmt.Println("No Product Available.")
		return
	}

	var productID int
	fmt.Scanln(&productID)

	var selectedProduct *model.Products
	for _, product := range products {
		if product.ID == productID {
			selectedProduct = &product
			break
		}
	}
	if selectedProduct == nil {
		fmt.Println("Product Not Found")
		return
	}

	fmt.Printf("You Selected Id : %d\n,Name : %s\n , Price : %.2f\n,Category : %s\n", selectedProduct.ID, selectedProduct.Name, selectedProduct.Price, selectedProduct.Category)

	//insert to table sales
	query := `Insert into sales(order_date, customer_id) values (?,?)`
	_, err := db.Exec(query, time.Now(), customerID)
	if err != nil {
		log.Fatalf("Failed to insert row : %v", err)
	}

	fmt.Println("Thank you for your purchase!")
}

// showProduct show list product
func showProduct(db *sql.DB) []model.Products {
	query := `select id,name,price,category from products`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var products []model.Products
	fmt.Println("Available Product")
	for rows.Next() {
		var product model.Products
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Category); err != nil {
			log.Fatalf("failed to scan product : %v", err)
		}
		fmt.Printf("ID : %d\n,Name : %s\n,Price : %f\n , category : %s\n", product.ID, product.Name, product.Price, product.Category)
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return products
}
