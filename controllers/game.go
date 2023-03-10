package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ropehapi/scrum-poker/database"
	"github.com/ropehapi/scrum-poker/models"
)

func GameIndex(w http.ResponseWriter, r *http.Request) {
	var games []models.Game
	database.DB.Find(&games)
	json.NewEncoder(w).Encode(games)	
}
