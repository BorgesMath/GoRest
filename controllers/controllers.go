package controllers

import (
	"APIGoRest/database"
	"APIGoRest/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {

	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)

}

func RetornarUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64) // O 10 é a base numerica utilizada

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, personalidade := range models.Personalidades {
		if personalidade.Id == int(id) { // Converta id para int para compararar
			json.NewEncoder(w).Encode(personalidade)
			return // Adicione um return para evitar loop
		}
	}

	http.Error(w, "Personalidade não encontrada", http.StatusNotFound) // Retorna um erro se a personalidade não for encontrada
}
