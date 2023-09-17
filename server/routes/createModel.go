package routes

import (
	"net/http"
	"server/controllers"
)

func createModel(w http.ResponseWriter, r *http.Request) {

	controllers.CreateDB()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Database created successfully."))

}
