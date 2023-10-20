package routes

import (
	controllers "GoRestApi/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func Init() {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id:[0-9]+}", controllers.PersonalidadePorId).Methods("GET")
	r.HandleFunc("/api/personalidades", controllers.NovaPersonalidade).Methods("POST")
	r.HandleFunc("/api/personalidades/{id:[0-9]+}", controllers.DeletaPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id:[0-9]+}", controllers.EditaPersonalidade).Methods("Put")

	http.ListenAndServe(":8000", r)
}
