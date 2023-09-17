package controllers

import (
	"fmt"
	"server/db"
)

func CreateTemporaryTable() {
	createTableQuery := `        
		CREATE TEMPORARY TABLE if not exists sistemadevotos.TempCargo (
		id INT,
		cargo VARCHAR(50)
	)
`
	var err error

	_, err = db.Db.Exec(createTableQuery)
	if err != nil {
		panic(err.Error())
	}

	loadCSVQuery := `
        LOAD DATA INFILE 'C:/ProgramData/MySQL/MySQL Server 8.0/Uploads/cargos.csv'
        INTO TABLE sistemadevotos.TempCargo
        FIELDS TERMINATED BY ','
        LINES TERMINATED BY '\n'
        IGNORE 1 ROWS  -- Skip header row if it exists
        (id, cargo);  -- Specify columns to load from CSV
    `

	_, err = db.Db.Exec(loadCSVQuery)
	if err != nil {
		panic(err.Error())
	}

	Load := `
	INSERT INTO sistemadevotos.cargo (id, cargo) SELECT id, cargo FROM sistemadevotos.TempCargo

`

	_, err = db.Db.Exec(Load)
	if err != nil {
		panic(err.Error())
	}

	// fmt.Print(ExecuteQuery("SHOW DATABASES"))

	fmt.Println("Table 'SistemaDeVotos' created successfully.")

}
