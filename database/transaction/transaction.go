package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// create sql.DB object
	database, error := sql.Open("mysql", "root:test@tcp(127.0.0.1)/testDatabase")
	if error != nil {
		log.Fatal(error)
	}
	defer database.Close()

	// start transaction
	transaction, error := database.Begin()
	if error != nil {
		log.Panic(error)
	}
	defer transaction.Rollback() // if error, rollback

	// excute INSERT
	_, error = transaction.Exec("INSERT INTO testTable (name) VALUES (?)", "transactionInsert_1")
	if error != nil {
		// print error message and call panic()
		// panic() run defer
		log.Panic(error)
	}

	_, error = transaction.Exec("INSERT INTO testTable (name) VALUES (?)", "transactionInsert_2")
	if error != nil {
		log.Panic(error)
	}

	// commit transaction
	error = transaction.Commit()
	if error != nil {
		log.Panic(error)
	}
}
