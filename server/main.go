package main

import (
	"net/http"

	"server/routes" // Import package

	"server/db"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// db.Db.Exec("SHOW DATABASES")
	defer db.Db.Close()
	// Use the SetupRoutes function from your routes package
	routes.SetupRoutes(r)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
