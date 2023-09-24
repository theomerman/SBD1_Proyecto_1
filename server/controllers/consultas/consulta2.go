package consultas

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"server/controllers/structures"
)

func GetConsult2(rows *sql.Rows, query int) ([]byte, error) {
	var vector []structures.Consulta2
	for rows.Next() {
		var column1Value string
		var column2Value int

		err := rows.Scan(&column1Value, &column2Value)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		vector = append(vector, structures.Consulta2{Partido: column1Value, Cantidad: column2Value})
	}
	jsonData, err := json.MarshalIndent(structures.Consult2{Consulta: query, Rows: len(vector), Return: vector}, "", "  ")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return jsonData, nil
}
