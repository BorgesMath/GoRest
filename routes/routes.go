package routes

import (
	ct "APIGoRest/controllers"
	"APIGoRest/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", ct.Home)
	r.HandleFunc("/api/personalidades", ct.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", ct.RetornarUmaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", ct.CriaUmaNovaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", ct.DeletaUmaPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", ct.EditarUmaPersonalidade).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
