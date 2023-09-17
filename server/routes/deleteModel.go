package routes

import (
	"fmt"
	"net/http"
	"server/controllers"
)

func deleteModel(w http.ResponseWriter, r *http.Request) {

	controllers.DeleteDB()
	w.WriteHeader(http.StatusOK)
	fmt.Println("Database deleted successfully.")
	w.Write([]byte("Database deleted successfully."))

}
