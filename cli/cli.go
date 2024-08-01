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
	user, customerID := Login(db)

	if user == "admin" {
		fmt.Println("ADMIN MENU!!!")
		AdminMenu(db)

	} else if user == "customer" {
		fmt.Println("CUSTOMER MENU!!!")
		CustomerMenu(db, customerID)
	} else {
		fmt.Println("Goodbye!")
		os.Exit(1)

	}

}

func Login(db *sql.DB) (string, int) {
	for {

		fmt.Println("\nSelect Menu:")
		fmt.Println("1. Login")
		fmt.Println("2. Sign Up")
		fmt.Println("99. Exit")

		fmt.Printf("\nEnter the number of the menu you want to access: ")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			result, customerID := handler.Login(db)
			if result != "invalid" {
				return result, customerID
			} else {
				fmt.Println("Invalid email or password")
			}
		case 2:
			handler.SignUp(db)
		case 99:
			fmt.Println("Exit")
			return "exit", 0
		default:
			fmt.Println("Please input valid number")
		}

	}
}

func AdminMenu(db *sql.DB) {
	for {
		fmt.Println("\nSelect Menu:")
		fmt.Println("1. Add Product")
		fmt.Println("2. Add Employee")
		fmt.Println("3. Order Reports")
		fmt.Println("4. Stock Reports")
<<<<<<< HEAD
		fmt.Println("5. User Reports")
		fmt.Println("0. Exit")
=======
		fmt.Println("5. Delete Product")
		fmt.Println("99. Exit")
>>>>>>> 560bfb7da77df89f0aea19706072d4c262f02f8c
		fmt.Printf("\nEnter the number of the menu you want to access: ")

		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			handler.AddProduct(db)
		case 2:
			handler.AddEmployee(db)
		case 3:
			handler.OrderReports(db)
		case 4:
			handler.StockReports(db)
		case 5:
<<<<<<< HEAD
			handler.UserReports(db)
		case 0:
=======
			handler.DeleteProduct(db)
		case 99:
>>>>>>> 560bfb7da77df89f0aea19706072d4c262f02f8c
			fmt.Println("Exit")
			return
		default:
			fmt.Println("Please input valid number")
		}
	}
}

func CustomerMenu(db *sql.DB, customerId int) {
	for {
		fmt.Println("\nSelect Menu:")
		fmt.Println("1. BuyProduct")
		fmt.Println("99. Exit")
		fmt.Printf("\nEnter the number of the menu you want to access: ")

		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			handler.BuyProduct(db, customerId)
		case 99:
			fmt.Println("Exit")
			return
		default:
			fmt.Println("Please input valid number")
		}
	}

}
