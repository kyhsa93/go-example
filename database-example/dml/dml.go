package main

import (
	"database/sql"
	"fmt"
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

	// excute INSERT
	result, error := database.Exec("INSERT INTO testTable (name) VALUES (?)", "insert")
	if error != nil {
		log.Fatal(error)
	}

	// check sql.Result.RowsAffected()
	number, error := result.RowsAffected()
	if number == 1 {
		fmt.Println(number)
	}

	// create prepared statement
	statement, error := database.Prepare("UPDATE testTable SET name = ? WHERE id = ?")
	checkError(error)
	defer statement.Close()

	// excute prepared statement
	_, error = statement.Exec("update", 1)
	checkError(error)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
