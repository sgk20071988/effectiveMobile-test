package handlers

import (
	"effectiveMobileTest/internal/db"
	"encoding/json"
	"net/http"
)

func GetCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(db.Get())
}
