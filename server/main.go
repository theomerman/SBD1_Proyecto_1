package main

import (
	"net/http"
	"server/db"
	"server/routes"

	"github.com/gorilla/mux"
)

// Import package

func main() {
	r := mux.NewRouter()

	// db.Db.Exec("SHOW DATABASES")
	defer db.Db.Close()
	// Use the SetupRoutes function from your routes package
	routes.SetupRoutes(r)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)

}
