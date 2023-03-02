package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ropehapi/scrum-poker/controllers"
)

func HandleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/players", controllers.Index).Methods("GET")
	r.HandleFunc("/player", controllers.Store).Methods("POST")
	
	log.Fatal(http.ListenAndServe(":8000", r))
}
