package handlers

import (
	"encoding/json"
	"gocurr/curr_errors"
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
		WriteJSONError(w, http.StatusInternalServerError, "Problem with getting currencies")
		return
	}

	data, err := json.Marshal(currencies)
	if err != nil {
		log.Printf("JSON marshal error: %v", err)
		WriteJSONError(w, http.StatusInternalServerError, "Encoding error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (curHandler *CurrencyHandler) GetByCode(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")

	currency, err := curHandler.service.GetByCode(code)
	if err != nil {
		switch e := err.(type) {
		case *curr_errors.NotValidError:
			log.Printf("Invalid currency code: %s", code)
			WriteJSONError(w, http.StatusBadRequest, "Некорректный код валюты")
			return

		case *curr_errors.NotFoundError:
			log.Printf("Currency not found: code=%s", code)
			WriteJSONError(w, http.StatusNotFound, "Валюта не найдена")
			return

		default:
			log.Printf("Unexpected error in GetByCode: %v", e)
			WriteJSONError(w, http.StatusInternalServerError, "Внутренняя ошибка сервера")
			return
		}
	}

	data, err := json.Marshal(currency)
	if err != nil {
		log.Printf("JSON marshal error: %v", err)
		WriteJSONError(w, http.StatusInternalServerError, "Ошибка кодирования ответа")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (curHandler *CurrencyHandler) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Printf("ParseForm error: %v", err)
		WriteJSONError(w, http.StatusBadRequest, "Некорректные данные формы")
		return
	}

	name := r.FormValue("name")
	code := r.FormValue("code")
	sign := r.FormValue("sign")

	currency, err := curHandler.service.CreateCurrency(name, code, sign)
	if err != nil {
		switch e := err.(type) {
		case *curr_errors.NotValidError:
			log.Printf("Invalid currency data: %s", code)
			WriteJSONError(w, http.StatusBadRequest, "Некорректные данные валюты")
			return

		case *curr_errors.CurrencyAlreadyExistsError:
			log.Printf("Currency is already exists: code=%s", code)
			WriteJSONError(w, http.StatusConflict, "Валюта уже сущесвует")
			return

		default:
			log.Printf("Unexpected error in GetByCode: %v", e)
			WriteJSONError(w, http.StatusInternalServerError, "Внутренняя ошибка сервера")
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(currency)
}
