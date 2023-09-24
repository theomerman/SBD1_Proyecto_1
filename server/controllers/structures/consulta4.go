package structures

type Consult4 struct {
	Consulta int         `json:"consulta"`
	Rows     int         `json:"rows"`
	Return   []Consulta4 `json:"return"`
}

type Consulta4 struct {
	Cantidad int    `json:"cantidad"`
	Partido  string `json:"partido"`
}
