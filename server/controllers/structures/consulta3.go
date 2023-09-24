package structures

type Consult3 struct {
	Consulta int         `json:"consulta"`
	Rows     int         `json:"rows"`
	Return   []Consulta3 `json:"return"`
}

type Consulta3 struct {
	Nombre  string `json:"nombre"`
	Partido string `json:"partido"`
}
