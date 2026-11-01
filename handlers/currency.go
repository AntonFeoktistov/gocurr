package handlers

import (
	"encoding/json"
	"gocurr/service"
	"log"
	"net/http"
)

type CurrencyHandler struct {
	service *service.CurrencyService
}

func NewCurrencyHandler(service *service.CurrencyService) *CurrencyHandler {
	return &CurrencyHandler{service: service}
}
func (curHandler *CurrencyHandler) List(w http.ResponseWriter, r *http.Request) {
	currencies, err := curHandler.service.ListCurrencies()
	if err != nil {
		log.Printf("ListCurrencies error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"Problem with getting currencies"}`))
		return
	}

	data, err := json.Marshal(currencies)
	if err != nil {
		log.Printf("JSON marshal error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"encoding error"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
