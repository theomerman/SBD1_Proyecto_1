package consultas

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"server/controllers/structures"
)

func GetConsult7(rows *sql.Rows, query int) ([]byte, error) {
	var vector []structures.Consulta7
	for rows.Next() {
		var column1Value int
		var column2Value int

		err := rows.Scan(&column1Value, &column2Value)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		vector = append(vector, structures.Consulta7{Edad: column1Value, Cantidad: column2Value})
	}
	jsonData, err := json.MarshalIndent(structures.Consult7{Consulta: query, Rows: len(vector), Return: vector}, "", "  ")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return jsonData, nil
}
