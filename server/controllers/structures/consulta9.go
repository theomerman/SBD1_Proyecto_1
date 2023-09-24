package structures

type Consult9 struct {
	Consulta int         `json:"consulta"`
	Rows     int         `json:"rows"`
	Return   []Consulta9 `json:"return"`
}

type Consulta9 struct {
	Mesa   int    `json:"mesa"`
	Votos  int    `json:"votos"`
	Nombre string `json:"nombre"`
}