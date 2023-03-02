package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequests() {
	r := mux.NewRouter()

	log.Fatal(http.ListenAndServe(":8000", r))
}
