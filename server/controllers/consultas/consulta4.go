package consultas

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"server/controllers/structures"
)

func GetConsult4(rows *sql.Rows, query int) ([]byte, error) {
	var vector []structures.Consulta4
	for rows.Next() {
		var column1Value int
		var column2Value string

		err := rows.Scan(&column1Value, &column2Value)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		vector = append(vector, structures.Consulta4{Cantidad: column1Value, Partido: column2Value})
	}
	jsonData, err := json.MarshalIndent(structures.Consult4{Consulta: query, Rows: len(vector), Return: vector}, "", "  ")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return jsonData, nil
}
