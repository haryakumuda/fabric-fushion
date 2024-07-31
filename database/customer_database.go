package database

import (
	"database/sql"
	"encoding/json"
	"fabric-fushion/helper"
	"log"
	"time"
)

type Database struct {
	DB *sql.DB
}

func (db *Database) InsertSale(CustomerID int, products map[int]int) error {
	//convert map to json
	productsJSON, err := json.Marshal(helper.ConvertToJSON(products))
	if err != nil {
		return err
	}

	// call store procedure
	query := `Call InsertSale(?,?,?)`
	_, err = db.DB.Exec(query, CustomerID, time.Now(), productsJSON)
	if err != nil {
		return err
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
