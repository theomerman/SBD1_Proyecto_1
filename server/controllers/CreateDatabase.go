package controllers

import (
	"fmt"

	"server/db"

	_ "github.com/go-sql-driver/mysql"
)

var dbName = "sistemadevotos"

func CreateDB() {

	// Create a new database
	var err error
	_, err = db.Db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		panic(err.Error())
	}
	// Use the database
	CreateTables()
	// CreateTemporaryTable()
	fmt.Printf("Database '%s' created successfully.\n", dbName)

}

func CreateTables() {
	var err error

	for i := 0; i < len(db.QueryVector); i++ {
		_, err = db.Db.Exec(db.QueryVector[i])
		if err != nil {
			panic(err.Error())
		}
	}

	// fmt.Print(ExecuteQuery("SHOW DATABASES"))

	// fmt.Println("Table 'SistemaDeVotos' created successfully.")

}

func ExecuteQuery(query string) string {
	// Execute the "SHOW DATABASES" query
	rows, err := db.Db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	// Iterate through the result set and print database names
	var dbName string
	var text string
	fmt.Println("List of databases:")
	for rows.Next() {
		if err := rows.Scan(&dbName); err != nil {
			panic(err.Error())
		}
		text += dbName + "\n"
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		panic(err.Error())
	}
	return text
}