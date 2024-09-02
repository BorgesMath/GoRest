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
	w.Header().Set("Content-type", "application/json")
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

	var NovaPersonalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&NovaPersonalidade)
	database.DB.Create(&NovaPersonalidade)
	json.NewEncoder(w).Encode(NovaPersonalidade)

}

func DeletaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64) // O 10 é a base numerica utilizada

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		errorMessage := fmt.Sprintf("Erro transformar Id: %v", err)
		fmt.Println(errorMessage)

		return
	}

	var p models.Personalidade
	database.DB.Delete(&p, id)
	json.NewEncoder(w).Encode(p)

}

func EditarUmaPersonalidade(w http.ResponseWriter, r *http.Request) {

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
	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Save(&p)
	json.NewEncoder(w).Encode(p)
}
