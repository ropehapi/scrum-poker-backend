package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ropehapi/scrum-poker/database"
	"github.com/ropehapi/scrum-poker/models"
)

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
