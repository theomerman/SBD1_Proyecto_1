// routes/routes.go
package routes

import (
	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	// Define your API routes
	r.HandleFunc("/crearmodelo", createModel).Methods("GET")
	r.HandleFunc("/eliminarmodelo", deleteModel).Methods("GET")
	r.HandleFunc("/cargartabtemp", loadDB).Methods("GET")
	r.HandleFunc("/consulta1", consulta1).Methods("GET")
	r.HandleFunc("/consulta2", consulta2).Methods("GET")
	r.HandleFunc("/consulta3", consulta3).Methods("GET")
	r.HandleFunc("/consulta4", consulta4).Methods("GET")
	r.HandleFunc("/consulta5", consulta5).Methods("GET")
	r.HandleFunc("/consulta6", consulta6).Methods("GET")
	r.HandleFunc("/consulta7", consulta7).Methods("GET")
	r.HandleFunc("/consulta8", consulta8).Methods("GET")
	r.HandleFunc("/consulta9", consulta9).Methods("GET")
	r.HandleFunc("/consulta10", consulta10).Methods("GET")
	r.HandleFunc("/consulta11", consulta11).Methods("GET")

}
