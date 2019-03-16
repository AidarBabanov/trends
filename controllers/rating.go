package controllers

import (
	"encoding/json"
	"net/http"
	"time"
	"trends/database"
	"trends/models"
)

func create(w http.ResponseWriter, r *http.Request) {
	var trends []models.Trends
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&trends); err != nil {
		respondWithError(w, http.StatusBadRequest, "Wrong data format")
		return
	}

	now := time.Now()
	db := database.DB
	for _, trend := range trends {
		trend.TrackedAt = &now
		db.Create(trend)
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "OK"})
}

func get(w http.ResponseWriter, r *http.Request) {
	db:=database.DB
	var trends []models.Trends
	db.Find(&trends).Order("created_at").Order("value")
}
