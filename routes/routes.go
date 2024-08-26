package routes

import (
	ct "APIGoRest/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", ct.Home)
	r.HandleFunc("/api/personalidades", ct.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", ct.RetornarUmaPersonalidade).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
