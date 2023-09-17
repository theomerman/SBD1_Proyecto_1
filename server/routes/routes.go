// routes/routes.go
package routes

import (
	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	// Define your API routes
	r.HandleFunc("/", createModel).Methods("GET")
	r.HandleFunc("/eliminarmodelo", deleteModel).Methods("GET")
}
