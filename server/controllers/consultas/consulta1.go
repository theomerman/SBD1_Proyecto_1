package consultas

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"server/controllers/structures"
)

func GetConsult1(rows *sql.Rows, query int) ([]byte, error) {
	var vector []structures.Consulta1
	for rows.Next() {
		var column1Value string
		var column2Value string
		var column3Value string

		err := rows.Scan(&column1Value, &column2Value, &column3Value)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		vector = append(vector, structures.Consulta1{Presidente: column1Value, Vicepresidente: column2Value, Partido: column3Value})
	}
	jsonData, err := json.MarshalIndent(structures.Consult1{Consulta: query, Rows: len(vector), Return: vector}, "", "  ")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return jsonData, nil
}
