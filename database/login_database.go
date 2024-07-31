package database

import (
	"database/sql"
	"fabric-fushion/model"
	"log"
)

func AddUser(db *sql.DB, user model.User) (int64, error) {
	query := `INSERT INTO users (email, password, role_id) VALUES (?, ?, ?)`
	result, err := db.Exec(query, user.Email, user.Password, user.RoleId)
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
	query := `INSERT INTO customers (user_id, name, phone_number) VALUES (?, ?, ?)`
	result, err := db.Exec(query, customer.UserId, customer.Name, customer.PhoneNumber)
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

func GetUser(db *sql.DB, user model.UserLogin) (bool, int, int) {
	var id, roleId int
	var dbEmail, dbPassword string

	// Query to find a user with the given email and password
	query := `SELECT id, email, password, role_id FROM users WHERE email = ? AND password = ?`
	err := db.QueryRow(query, user.Email, user.Password).Scan(&id, &dbEmail, &dbPassword, &roleId)
	if err != nil {
		if err == sql.ErrNoRows {
			// No matching user found
			return false, 0, 0
		}
		log.Fatal("Error executing query: ", err)
	}

	// If we reach here, a matching user was found
	return true, roleId, id
}

func AddEmployee(db *sql.DB, employee model.Employee) (int64, error) {
	query := `INSERT INTO employees (user_id, name, position_id) VALUES (?, ?, ?)`
	result, err := db.Exec(query, employee.UserId, employee.Name, employee.PositionId)
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
