package controllers

import (
	"encoding/json"
	"net/http"
	"time"
	"trends/database"
	"trends/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	database.DB.AutoMigrate(models.Trends{})

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

func Get(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	var trends []models.Trends
	db.Find(&trends).Order("created_at").Order("value")
	respondWithJSON(w, http.StatusOK, trends)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := database.DB
	db.DropTableIfExists(models.Trends{})
	respondWithJSON(w, http.StatusOK, map[string]string{"message": "OK"})
}
