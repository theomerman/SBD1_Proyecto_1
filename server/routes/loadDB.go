package routes

import (
	"net/http"
	"server/controllers"
)

func loadDB(w http.ResponseWriter, r *http.Request) {
	controllers.TemporaryTable()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Temporary table loaded successfully into main Database."))
}
