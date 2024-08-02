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
		fmt.Println("0. Exit")

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
		case 0:
			fmt.Println("Exit")
			return "exit", 0
		}
	}
}

func AdminMenu(db *sql.DB) {
	for {
		fmt.Println("\nSelect Menu:")
		fmt.Println("1. Order Reports")
		fmt.Println("2. Stock Reports")
		fmt.Println("3. User Reports")
		fmt.Println("4. Add Product")
		fmt.Println("5. Add Employee")
		fmt.Println("6. Update Stock")
		fmt.Println("7. Delete Product")
		fmt.Println("8. Delete Customer")
		fmt.Println("9. Delete Admin")
		fmt.Println("99. Exit")
		fmt.Printf("\nEnter the number of the menu you want to access: ")

		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			handler.OrderReports(db)
		case 2:
			handler.StockReports(db)
		case 3:
			handler.UserReports(db)
		case 4:
			handler.AddProduct(db)
		case 5:
			handler.AddEmployee(db)
		case 6:
			handler.UpdateStock(db)
		case 7:
			handler.DeleteProduct(db)
		case 8:
			handler.DeleteCustomer(db)
		case 9:
			handler.DeleteAdmin(db)
		case 99:
			fmt.Println("Exit")
			return
		}
	}
}

func CustomerMenu(db *sql.DB, customerId int) {
	for {
		fmt.Println("\nSelect Menu:")
		fmt.Println("1. ShowProduct")
		fmt.Println("2. BuyProduct")
		fmt.Println("3. OrderHistory")
		fmt.Println("99. Exit")
		fmt.Printf("\nEnter the number of the menu you want to access: ")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println(err)
		}

		switch choice {
		case 1:
			handler.ShowProduct(db)
		case 2:
			handler.BuyProduct(db, customerId)
		case 3:
			handler.OrderHistory(db, customerId)
		case 99:
			fmt.Println("Exit")
			return
		default:
			fmt.Println("invalid number , please try again")
		}
	}

}
