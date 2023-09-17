package controllers

import (
	"server/db"

	_ "github.com/go-sql-driver/mysql"
)

func DeleteDB() {

	db.Db.Exec("DROP DATABASE IF EXISTS " + dbName)

}
