package structures

type Consult8 struct {
	Consulta int         `json:"consulta"`
	Rows     int         `json:"rows"`
	Return   []Consulta8 `json:"return"`
}

type Consulta8 struct {
	Presidente     string `json:"presidente"`
	Vicepresidente string `json:"vicepresidente"`
	Votos          int    `json:"votos"`
}