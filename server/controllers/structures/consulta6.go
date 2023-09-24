package structures

type Consult6 struct {
	Consulta int         `json:"consulta"`
	Rows     int         `json:"rows"`
	Return   []Consulta6 `json:"return"`
}


type Consulta6 struct {
	Votos    string `json:"voto"`
	Cantidad int    `json:"cantidad"`
}