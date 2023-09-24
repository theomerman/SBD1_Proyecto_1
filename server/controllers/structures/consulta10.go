package structures

type Consult10 struct {
	Consulta int         `json:"consulta"`
	Rows     int         `json:"rows"`
	Return   []Consulta10 `json:"return"`
}

type Consulta10 struct {
	Hora  string `json:"hora"`
	Votos int    `json:"votos"`
}