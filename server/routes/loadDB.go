package routes

import (
	"net/http"
	"server/controllers"
)

func loadDB(w http.ResponseWriter, r *http.Request) {
	_, err := controllers.TemporaryTable()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error al cargar datos a la base de datos."))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Tabla temporal cargada con éxito."))
}
