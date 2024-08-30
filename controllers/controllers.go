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
		errorMessage := fmt.Sprintf("Erro transformar Id: %v", err)
		fmt.Println(errorMessage)

		return
	}

	var p models.Personalidade

	database.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)

}

func CriaUmaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {

	var NovaP models.Personalidade
	json.NewDecoder(r.Body).Decode(&NovaP)
	database.DB.Create(&NovaP)
	json.NewEncoder(w).Encode(NovaP)

}
