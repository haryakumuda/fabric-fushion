package database

import (
	"database/sql"
	"fabric-fushion/model"
	"log"
)

func AddUser(db *sql.DB, user model.User) (int64, error) {
	query := `INSERT INTO users (email, password, role) VALUES (?, ?, ?)`
	result, err := db.Exec(query, user.Email, user.Password, user.Role)
	if err != nil {
		return 0, err
	}

	// Get the ID of the newly inserted row
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	log.Println("User added successfully with ID:", id)
	return id, nil
}

func AddCustomer(db *sql.DB, customer model.Customer) (int64, error) {
	query := `INSERT INTO customers (user_id, email, name, phone_number) VALUES (?, ?, ?, ?)`
	result, err := db.Exec(query, customer.UserId, customer.Email, customer.Name, customer.PhoneNumber)
	if err != nil {
		return 0, err
	}

	// Get the ID of the newly inserted row
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	log.Println("Customer added successfully with ID:", id)
	return id, nil
}

func GetUser(db *sql.DB, user model.UserLogin) (bool, string, int) {
	var id int
	var dbEmail, dbPassword, role string

	// Query to find a user with the given email and password
	query := `SELECT id, email, password, role FROM users WHERE email = ? AND password = ?`
	err := db.QueryRow(query, user.Email, user.Password).Scan(&id, &dbEmail, &dbPassword, &role)
	if err != nil {
		if err == sql.ErrNoRows {
			// No matching user found
			return false, "", 0
		}
		log.Fatal("Error executing query: ", err)
	}

	// If we reach here, a matching user was found
	return true, role, id
}

func AddEmployee(db *sql.DB, employee model.Employee) (int64, error) {
	query := `INSERT INTO employees (user_id, email, name, position) VALUES (?, ?, ?, ?)`
	result, err := db.Exec(query, employee.UserId, employee.Email, employee.Name, employee.Position)
	if err != nil {
		return 0, err
	}

	// Get the ID of the newly inserted row
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	log.Println("Customer added successfully with ID:", id)
	return id, nil
}
