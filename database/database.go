package database

import (
	"database/sql"
	"fabric-fushion/config"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // Make sure to import the MySQL driver
)

var conn *sql.DB

// GenerateDatabase initializes the global database connection.
func GenerateDatabase() {
	var err error
	conn, err = sql.Open("mysql", config.DatabaseConfig())
	if err != nil {
		log.Fatal("Error Loading Database: ", err)
	}
	fmt.Println("Database Connected")

	err = conn.Ping()
	if err != nil {
		log.Fatal("Error Ping Database: ", err)
	}
	fmt.Println("Ping Database Done")
}

// GetDatabase returns the global database connection.
func GetDatabase() *sql.DB {
	return conn
}
