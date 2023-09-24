package controllers

import (
	"fmt"
	"server/db"
)

func TemporaryTable() (string, error) {
	x := db.LoadTemporal()

	for i := 0; i < len(x); i++ {

		_, err := db.Db.Exec(x[i])
		if err != nil {
			fmt.Println(err.Error())
			return "", err
		}
	}
	fmt.Println("Temporary table loaded successfully into main Database.")
	return "Temporary table loaded successfully into main Database.", nil

}
