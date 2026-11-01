package handlers

import (
	"encoding/json"
	"net/http"
)

func CurrencyHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(name); err != nil {
		http.Error(w, "encoding error", http.StatusInternalServerError)
	}
}
