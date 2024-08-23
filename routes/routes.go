package routes

import (
	ct "APIGoRest/controllers"
	"log"
	"net/http"
)

func HandleRequest() {
	http.HandleFunc("/", ct.Home)
	http.HandleFunc("/api/personalidades", ct.TodasPersonalidades)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
