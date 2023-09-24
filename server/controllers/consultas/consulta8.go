package consultas

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"server/controllers/structures"
)

func GetConsult8(rows *sql.Rows, query int) ([]byte, error) {
	var vector []structures.Consulta8
	for rows.Next() {
		var column1Value string
		var column2Value string
		var column3Value int

		err := rows.Scan(&column1Value, &column2Value, &column3Value)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		vector = append(vector, structures.Consulta8{Presidente: column1Value, Vicepresidente: column2Value, Votos: column3Value})
	}
	jsonData, err := json.MarshalIndent(structures.Consult8{Consulta: query, Rows: len(vector), Return: vector}, "", "  ")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return jsonData, nil
}
