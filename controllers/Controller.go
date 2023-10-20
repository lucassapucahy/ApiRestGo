package controllers

import (
	data "GoRestApi/Data"
	"GoRestApi/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "home page")
	w.WriteHeader(http.StatusOK)
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	p := data.BuscarPersonalidades()
	json.NewEncoder(w).Encode(p)
	w.WriteHeader(http.StatusOK)
}

func PersonalidadePorId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	p := data.BuscarPersonalidadePorId(id)

	json.NewEncoder(w).Encode(p)
	w.WriteHeader(http.StatusOK)
}

func NovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var personalidade models.Personalidade

	err := json.NewDecoder(r.Body).Decode(&personalidade)

	if err != nil {
		log.Fatal(err.Error())
	}

	personalidadeInserida := data.InserirPersonalidade(personalidade)

	json.NewEncoder(w).Encode(personalidadeInserida)
	w.WriteHeader(http.StatusCreated)
}

func DeletaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	data.DeletaPersonalidade(id)

	w.WriteHeader(http.StatusNoContent)
}

func EditaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personalidade
	json.NewDecoder(r.Body).Decode(&p)
	idConvertido, _ := strconv.Atoi(id)
	p.Id = idConvertido
	data.EditaPersonalidade(p)
	json.NewEncoder(w).Encode(p)
	w.WriteHeader(http.StatusOK)
}
