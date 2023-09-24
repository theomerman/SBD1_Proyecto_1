package structures

type Consult7 struct {
	Consulta int         `json:"consulta"`
	Rows     int         `json:"rows"`
	Return   []Consulta7 `json:"return"`
}

type Consulta7 struct {
	Edad     int `json:"edad"`
	Cantidad int `json:"cantidad"`
}
