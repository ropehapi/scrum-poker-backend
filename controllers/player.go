package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ropehapi/scrum-poker/database"
	"github.com/ropehapi/scrum-poker/models"
)

func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var player models.Player
	database.DB.First(&player, id)
	json.NewEncoder(w).Encode(player)
}

func Index(w http.ResponseWriter, r *http.Request) {
	var players []models.Player
	database.DB.Find(&players)
	json.NewEncoder(w).Encode(players)
}

func Store(w http.ResponseWriter, r *http.Request) {
	var player models.Player
	json.NewDecoder(r.Body).Decode(&player)
	database.DB.Create(&player)
	json.NewEncoder(w).Encode(player)
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var player models.Player
	database.DB.First(&player, id)
	json.NewDecoder(r.Body).Decode(&player)
	database.DB.Save(&player)
	json.NewEncoder(w).Encode(player)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var player models.Player
	database.DB.Delete(&player, id)
	json.NewEncoder(w).Encode(player)
}
