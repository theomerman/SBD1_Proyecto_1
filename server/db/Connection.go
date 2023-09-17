package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Db = ConnectDB()

func ConnectDB() *sql.DB {

	// Define MySQL connection parameters
	dbUser := "root"
	dbPass := "iaKkC$#rcRe3eBdMqE?J"
	dbHost := "localhost"
	dbPort := "3306"

	// Create a database connection string
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/", dbUser, dbPass, dbHost, dbPort)

	// Connect to MySQL server
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	// defer db.Close()
	return db

}
