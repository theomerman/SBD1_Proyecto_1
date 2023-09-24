package consultas

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"server/controllers/structures"
)

func GetConsult9(rows *sql.Rows, query int) ([]byte, error) {
	var vector []structures.Consulta9
	for rows.Next() {
		var column1Value int
		var column2Value int
		var column3Value string

		err := rows.Scan(&column1Value, &column2Value, &column3Value)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		vector = append(vector, structures.Consulta9{Mesa: column1Value, Votos: column2Value, Nombre: column3Value})
	}
	jsonData, err := json.MarshalIndent(structures.Consult9{Consulta: query, Rows: len(vector), Return: vector}, "", "  ")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return jsonData, nil
}
