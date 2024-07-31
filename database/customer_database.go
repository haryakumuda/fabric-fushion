package database

import (
	"database/sql"
	"log"
	"time"
)

type Database struct {
	DB *sql.DB
}

func (db *Database) InsertSale(CustomerID int) error {
	query := `Insert into sales(order_date, customer_id) values (?,?)`
	_, err := db.DB.Exec(query, time.Now(), CustomerID)
	if err != nil {
		log.Fatalf("Failed to insert row : %v", err)
	}
	return nil
}

func (db *Database) ShowProducts() (*sql.Rows, error) {
	rows, err := db.DB.Query("CALL GetProductsWithCategory()")
	if err != nil {
		log.Fatal(err)
	}
	return rows, nil
}
