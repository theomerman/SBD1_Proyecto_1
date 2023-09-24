package routes

import (
	"fmt"
	"net/http"
	"server/controllers"
)

func consulta11(w http.ResponseWriter, r *http.Request) {

	JSON, err := controllers.ExecuteConsult(11)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error 500 - No se encontro la base de datos"))
		return
	}
	// fmt.Print(string(JSON))
	fmt.Println("status: 200 OK, consulta11")

	// w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(JSON)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error 500 - No se encontro la base de datos"))
		return
	}
}
