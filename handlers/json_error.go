package handlers

import (
	"encoding/json"
	"net/http"
)

func WriteJSONError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := map[string]string{
		"message": message,
	}

	data, err := json.Marshal(resp)
	if err != nil {
		w.Write([]byte(`{"message":"internal error"}`))
		return
	}
	w.Write(data)
}
