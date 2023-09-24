package structures

type Consult2 struct {
	Consulta int         `json:"consulta"`
	Rows     int         `json:"rows"`
	Return   []Consulta2 `json:"return"`
}
type Consulta2 struct {
	Partido  string `json:"partido"`
	Cantidad int    `json:"cantidad"`
}