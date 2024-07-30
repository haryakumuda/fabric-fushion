package cli

import (
	"database/sql"
	"fabric-fushion/handler"
	"fmt"
	"os"
)

func RunCLI(db *sql.DB) {

	handler.Welcome()

	// SignUp
	user := Login(db)

	if user == "admin" {
		// AdminMenu()
		fmt.Println("ADMIN MENU!!!")

	} else if user == "customer" {
		// CustomerMenu()
		fmt.Println("CUSTOMER MENU!!!")
	} else {

		fmt.Println("Goodbye!")
		os.Exit(1)

	}

}

func Login(db *sql.DB) string {
	for {

		fmt.Println("\nSelect Menu:")
		fmt.Println("1. Login")
		fmt.Println("2. Sign Up")
		fmt.Println("0. Exit")

		fmt.Printf("\nEnter the number of the menu you want to access: ")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			result := handler.Login(db)
			if result != "invalid" {
				return result
			} else {
				fmt.Println("Invalid email or password")
			}
		case 2:
			handler.SignUp(db)
		case 0:
			fmt.Println("Exit")
			return "exit"
		}
	}
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
