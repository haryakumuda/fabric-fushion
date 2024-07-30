package main

import (
	"database/sql"
	"fabric-fushion/cli"
	"fabric-fushion/config"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	fmt.Println("Start Application")

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error Loading Configuration")
	// }
	// fmt.Println("Config Done")

	db, err := sql.Open("mysql", config.DatabaseConfig())
	if err != nil {
		log.Fatal("Error Loading Database: ", err)
	}
	fmt.Println("Database Done")

	err = db.Ping()
	if err != nil {
		log.Fatal("Error Loading Database: ", err)
	}
	fmt.Println("Ping Database Done")

	defer db.Close()

	cli.RunCLI(db)
}
