package cli

import (
	"database/sql"
	"fmt"
)

func RunCLI(db *sql.DB) {

	// user := Login(db)

	// if user == "admin" {
	// 	AdminMenu()

	// } else if user == "customer" {
	// 	CustomerMenu()
	// }

}

func Login() string {

	// return antara "customer" / "admin"
	return "customer"
}

func AdminMenu() {
	for {
		fmt.Println("\nSelect Menu:")
		fmt.Println("1. Add Product")
		fmt.Println("2. Add Employee")
		fmt.Println("3. Order Reports")
		fmt.Println("4. Stock Reprots")
		fmt.Println("5. Exit")
		fmt.Printf("\nEnter the number of the menu you want to access: ")

		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			// handler.AddProduct()
		case 2:
			// handler.AddEmployee()
		case 3:
			// handler.OrderReports()
		case 4:
			// handler.StockReports()
		case 5:
			fmt.Println("Exit")
			return
		}
	}

}

func CustomerMenu() {

	for {
		fmt.Println("\nSelect Menu:")
		fmt.Println("1. BuyProduct")
		fmt.Println("5. Exit")
		fmt.Printf("\nEnter the number of the menu you want to access: ")

		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			// handler.BuyProduct()
		case 5:
			fmt.Println("Exit")
			return
		}
	}

}
