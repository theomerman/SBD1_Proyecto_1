package controllers

import (
	"fmt"
	"server/controllers/consultas"
	"server/db"

	_ "github.com/go-sql-driver/mysql"
)

func ExecuteConsult(query int) ([]byte, error) {

	_, err := db.Db.Exec("USE sistemadevotos")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	rows, err := db.Db.Query(db.GetQueryVector()[query-1])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	if query == 1 {
		return consultas.GetConsult1(rows, query)
	} else if query == 2 {
		return consultas.GetConsult2(rows, query)
	} else if query == 3 {
		return consultas.GetConsult3(rows, query)
	} else if query == 4 {
		return consultas.GetConsult4(rows, query)
	} else if query == 5 {
		return consultas.GetConsult5(rows, query)
	} else if query == 6 {
		return consultas.GetConsult6(rows, query)
	} else if query == 7 {
		return consultas.GetConsult7(rows, query)
	} else if query == 8 {
		return consultas.GetConsult8(rows, query)
	} else if query == 9 {
		return consultas.GetConsult9(rows, query)
	} else if query == 10 {
		return consultas.GetConsult10(rows, query)
	} else if query == 11 {
		return consultas.GetConsult11(rows, query)
	}
	err = rows.Err()
	return nil, err
}