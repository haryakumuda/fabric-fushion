package handler

import (
	"database/sql"
	"fabric-fushion/database"
	"fabric-fushion/model"
	"fmt"
	"strconv"
)

func AddProduct(db *sql.DB) {
	var product model.Product

	categories, err := database.ShowCategories(db)
	if err != nil {
		fmt.Println("Error showing categories: ", err)
		return
	}

	fmt.Println("Enter Product Name:")
	fmt.Scanln(&product.Name)

	var categoriesId []int64
	categoryMap := make(map[int64]string)
	for _, category := range categories {

		fmt.Printf("%d: %s\n", category.ID, category.Category)
		categoriesId = append(categoriesId, int64(category.ID))
		categoryMap[category.ID] = category.Category
	}
	fmt.Println("Enter Product Category ID:")
	for {
		var inputID int64
		fmt.Scanln(&inputID)

		if contains(categoriesId, inputID) {
			product.CategoryID = inputID
			break
		} else {
			fmt.Println("Please Input correct ID")
		}
	}
	fmt.Println("Enter Product Price:")
	for {
		var input string
		fmt.Scanln(&input)

		// Try to convert the input to float64
		price, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Please input a valid amount")
		} else {
			product.Price = price
			break
		}
	}

	fmt.Println("Enter Product Stock Quantity:")

	for {
		var input string
		fmt.Scanln(&input)

		// Try to convert the input to int64
		stock, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			fmt.Println("Please input a valid amount")
		} else {
			product.Stock = stock
			break
		}
	}

	fmt.Println("--- Product Summary ---")

	fmt.Println("Product Name: ", product.Name)
	fmt.Println("Product Price: ", product.Price)
	fmt.Println("Product Stock: ", product.Stock)
	fmt.Println("Product Category: ", categoryMap[product.CategoryID])
	fmt.Println("Are you sure want to add new product? (y/n)")
	var confirmation string
	for {
		fmt.Scanln(&confirmation)
		if confirmation == "y" || confirmation == "yes" {
			// Call the function
			_, err := database.AddProduct(db, product)
			if err != nil {
				fmt.Println("Error deleting product: ", err)
			} else {
				fmt.Println("Product added successfully!")
			}
			break
		} else if confirmation == "n" || confirmation == "no" {
			fmt.Println("Add product cancelled.")
			break
		} else {
			fmt.Printf("Please input valid input: ")
		}

	}

}

func DeleteProduct(db *sql.DB) {
	fmt.Println()

	products, err := database.ShowProducts(db)
	if err != nil {
		fmt.Println("Error showing products: ", err)
		return
	}
	// Create a map to store product details by ID
	productMap := make(map[int64]struct {
		Name string
	})

	var productsId []int64
	productsId = append(productsId, 0)
	for _, product := range products {
		fmt.Printf("%d: %s\n", product.ID, product.Name)
		productsId = append(productsId, int64(product.ID))
		productMap[int64(product.ID)] = struct {
			Name string
		}{Name: product.Name}
	}

	fmt.Printf("Enter the product ID you wish to delete: ")

	for {
		var inputID int64
		fmt.Scanln(&inputID)

		if contains(productsId, inputID) {
			productName := productMap[inputID].Name
			// Ask for user confirmation
			fmt.Printf("Are you sure you want to delete '%s'? (y/n): ", productName)

			var confirmation string
			for {
				fmt.Scanln(&confirmation)
				if confirmation == "y" || confirmation == "yes" {
					// Call the function
					_, err := database.DeleteProduct(db, inputID)
					if err != nil {
						fmt.Println("Error deleting product: ", err)
					} else {
						fmt.Println("Product deleted successfully.")
					}
					break
				} else if confirmation == "n" || confirmation == "no" {
					fmt.Println("Deletion cancelled.")
					break
				} else {
					fmt.Printf("Please input valid input: ")
				}

			}

		} else {
			fmt.Println("Please Input correct ID")
		}
	}

}

func DeleteCustomer(db *sql.DB) {
	fmt.Println()

	customers, err := database.ShowCustomer(db)
	if err != nil {
		fmt.Println("Error showing customer: ", err)
		return
	}
	// Create a map to store product details by ID
	customerMap := make(map[int64]model.ShowCustomer)

	var customersId []int64
	customersId = append(customersId, 0)
	for _, customer := range customers {
		fmt.Printf("%d: %s (%s)\n", customer.CustomerId, customer.Name, customer.Email)
		customersId = append(customersId, int64(customer.CustomerId))
		customerMap[int64(customer.CustomerId)] = customer
	}

	fmt.Printf("Enter the customer ID you wish to delete: ")

	for {
		var inputID int64
		fmt.Scanln(&inputID)

		if contains(customersId, inputID) {
			customer := customerMap[inputID]
			// Ask for user confirmation
			fmt.Printf("Are you sure you want to delete '%s' (%s)? (y/n): ", customer.Name, customer.Email)

			var confirmation string
			for {
				fmt.Scanln(&confirmation)
				if confirmation == "y" || confirmation == "yes" {
					// Call the function
					_, err := database.DeleteCustomer(db, customer.CustomerId)
					if err != nil {
						fmt.Println("Error deleting product: ", err)
						break
					}
					_, err = database.DeleteUser(db, customer.UserId)
					if err != nil {
						fmt.Println("Error deleting product: ", err)
					} else {
						fmt.Println("Product deleted successfully.")
					}
					break
				} else if confirmation == "n" || confirmation == "no" {
					fmt.Println("Deletion cancelled.")
					break
				} else {
					fmt.Printf("Please input valid input: ")
				}
			}
		} else {
			fmt.Println("Please Input correct ID")
		}
	}
}

// Helper function to check if an ID exists in a slice
func contains(slice []int64, id int64) bool {
	for _, v := range slice {
		if v == id {
			return true
		}
	}
	return false
}
