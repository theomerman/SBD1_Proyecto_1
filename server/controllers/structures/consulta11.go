package structures

type Consult11 struct {
	Consulta int          `json:"consulta"`
	Rows     int          `json:"rows"`
	Return   []Consulta11 `json:"return"`
}

type Consulta11 struct {
	Genero string `json:"genero"`
	Votos  int    `json:"votos"`
}
