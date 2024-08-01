package main

import (
	"fabric-fushion/cli"
	"fabric-fushion/database"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	fmt.Println("Start Application")
	database.GenerateDatabase()
	conn := database.GetDatabase()

	defer conn.Close()
	cli.RunCLI(conn)

}
