package structures

type Consult1 struct {
	Consulta int         `json:"consulta"`
	Rows     int         `json:"rows"`
	Return   []Consulta1 `json:"return"`
}

type Consulta1 struct {
	Presidente     string `json:"presidente"`
	Vicepresidente string `json:"vicepresidente"`
	Partido        string `json:"partido"`
}
