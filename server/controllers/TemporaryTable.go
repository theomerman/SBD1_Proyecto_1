package controllers

import (
	"fmt"
	"server/db"
)

func TemporaryTable() {
	x := db.LoadTemporal()

	for i := 0; i < len(x); i++ {

		_, err := db.Db.Exec(x[i])
		if err != nil {
			panic(err.Error())
		}
	}
	fmt.Println("Temporary table loaded successfully into main Database.")

}
