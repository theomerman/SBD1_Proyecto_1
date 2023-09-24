package consultas

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"server/controllers/structures"
)

func GetConsult3(rows *sql.Rows, query int) ([]byte, error) {
	var vector []structures.Consulta3
	for rows.Next() {
		var column1Value string
		var column2Value string

		err := rows.Scan(&column1Value, &column2Value)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		vector = append(vector, structures.Consulta3{Nombre: column1Value, Partido: column2Value})
	}
	jsonData, err := json.MarshalIndent(structures.Consult3{Consulta: query, Rows: len(vector), Return: vector}, "", "  ")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return jsonData, nil
}
