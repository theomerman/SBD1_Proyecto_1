package main

import (
	"net/http"
	"server/db"
	"server/routes"

	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
)

// Import package

func main() {
	r := mux.NewRouter()

	defer db.Db.Close()
	// Use the SetupRoutes function from your routes package
	routes.SetupRoutes(r)

	// Create a new CORS middleware handler
	corsHandler := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedOrigins([]string{"http://localhost:4200"}), // Add your Angular app's URL
	)

	http.Handle("/", corsHandler(r))
	http.ListenAndServe(":8080", nil)

}
