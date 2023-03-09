package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ropehapi/scrum-poker/controllers"
)

func HandleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/player/{id}", controllers.Show).Methods("GET")
	r.HandleFunc("/players", controllers.Index).Methods("GET")
	r.HandleFunc("/player", controllers.Store).Methods("POST")
	r.HandleFunc("/player/{id}", controllers.Update).Methods("PUT")
	r.HandleFunc("/player/{id}", controllers.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
