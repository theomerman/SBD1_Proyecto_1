package routes

import (
	"fmt"
	"net/http"
	"server/controllers"
)

func deleteModel(w http.ResponseWriter, r *http.Request) {

	controllers.DeleteDB()
	w.WriteHeader(http.StatusOK)
	fmt.Println("Modelo borrado éxitosamente.")
	w.Write([]byte("Modelo borrado éxitosamente."))

}
