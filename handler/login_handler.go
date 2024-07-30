package handler

import (
	"database/sql"
	"fabric-fushion/database"
	"fabric-fushion/model"
	"fmt"
	"strings"
)

func Welcome() {
	// Define the header text
	header := "WELCOME TO FABRIC FUSHION"
	// Calculate the length of the divider to match the length of the header
	headerLength := len(header)

	// Create a divider string with the same length as the header
	divider := strings.Repeat("=", headerLength)

	// Print the header with formatted divider
	fmt.Println()
	fmt.Println(divider)
	fmt.Println(header)
	fmt.Println(divider)
	fmt.Println()
}

func Login(db *sql.DB) string {
	fmt.Printf("Enter Your Email: ")
	var email string
	fmt.Scanln(&email)

	fmt.Printf("Enter Your Password: ")
	var password string
	fmt.Scanln(&password)

	userLogin := model.UserLogin{
		Email:    email,
		Password: password,
	}

	result, role := database.GetUser(db, userLogin)

	if result {
		return role
	} else {
		return "invalid"
	}

}

func SignUp(db *sql.DB) {
	fmt.Printf("Enter Your Name: ")
	var name string
	fmt.Scanln(&name)

	fmt.Printf("Enter Your Email: ")
	var email string
	fmt.Scanln(&email)

	fmt.Printf("Enter Your Password: ")
	var password string
	fmt.Scanln(&password)

	fmt.Printf("Enter Your Phone Number: ")
	var phoneNumber string
	fmt.Scanln(&phoneNumber)

	fmt.Println("\nAre you sure want to create user? (y/n)")
	fmt.Println("Name: ", name)
	fmt.Println("Email: ", email)
	fmt.Println("Password: ", password)
	fmt.Println("Phone Number: ", phoneNumber)
	for {
		var answer string
		fmt.Scanln(&answer)

		if answer == "y" {
			break
		} else if answer == "n" {
			return
		} else {
			fmt.Println("Please input valid answer! (y/n)")
		}

	}

	user := model.User{
		Id:       0,
		Email:    email,
		Password: password,
		Role:     "customer",
	}

	id, err := database.AddUser(db, user)

	if err != nil {
		fmt.Println("Error Creating User: ", err)
	}

	customer := model.Customer{
		Id:          0,
		UserId:      id,
		Email:       email,
		Name:        name,
		PhoneNumber: phoneNumber,
	}

	_, err = database.AddCustomer(db, customer)

	if err != nil {
		fmt.Println("Error Creating User: ", err)
	}

}

func AddEmployee(db *sql.DB) {
	fmt.Printf("Enter Admin Name: ")
	var name string
	fmt.Scanln(&name)

	fmt.Printf("Enter Admin Email: ")
	var email string
	fmt.Scanln(&email)

	fmt.Printf("Enter Admin Password: ")
	var password string
	fmt.Scanln(&password)

	fmt.Printf("Enter Admin Position: ")
	var position string
	fmt.Scanln(&position)

	fmt.Println("\nAre you sure want to create user? (y/n)")
	fmt.Println("Name: ", name)
	fmt.Println("Email: ", email)
	fmt.Println("Password: ", password)
	fmt.Println("Position: ", position)
	for {
		var answer string
		fmt.Scanln(&answer)

		if answer == "y" {
			break
		} else if answer == "n" {
			return
		} else {
			fmt.Println("Please input valid answer! (y/n)")
		}

	}

	user := model.User{
		Id:       0,
		Email:    email,
		Password: password,
		Role:     "admin",
	}

	id, err := database.AddUser(db, user)

	if err != nil {
		fmt.Println("Error Creating User: ", err)
	}

	employee := model.Employee{
		Id:       0,
		UserId:   id,
		Email:    email,
		Name:     name,
		Position: position,
	}

	_, err = database.AddEmployee(db, employee)

	if err != nil {
		fmt.Println("Error Creating User: ", err)
	}

}
