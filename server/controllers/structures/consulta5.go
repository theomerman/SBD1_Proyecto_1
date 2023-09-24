package structures

type Consult5 struct {
	Consulta int         `json:"consulta"`
	Rows     int         `json:"rows"`
	Return   []Consulta5 `json:"return"`
}

type Consulta5 struct {
	Departamento string `json:"departamento"`
	Voto         int    `json:"votos"`
}
