package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	database, error := sql.Open("mysql", "root:test@tcp(127.0.0.1)/testDatabase")
	if error != nil {
		log.Fatal(error)
	}
	defer database.Close()

	// single row sql query
	var name string
	error = database.QueryRow("SELECT name FROM testTable WHERE id = 1").Scan(&name)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(name)

	// multirow sql query
	var id int
	rows, error := database.Query("SELECT id FROM testTable WHERE id <= ?", 10)
	if error != nil {
		log.Fatal(error)
	}
	defer rows.Close() // rows must be closed

	for rows.Next() {
		error := rows.Scan(&id)
		if error != nil {
			log.Fatal(error)
		}
		fmt.Println(id)
	}
}
