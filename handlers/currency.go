package handlers

import (
	"encoding/json"
	"net/http"
)

func CurrencyHandler(w http.ResponseWriter, r *http.Request) {
	currencies := []string{}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(currencies); err != nil {
		http.Error(w, "encoding error", http.StatusInternalServerError)
	}
}
