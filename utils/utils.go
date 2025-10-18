package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJSON(r *http.Request, payload interface{}) error {
	if r.Body == nil {
		return fmt.Errorf("request body is empty")
	}
	
	defer r.Body.Close()
	
	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, statusCode int, v interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, statusCode int, err error) {
	WriteJSON(w, statusCode, map[string]string{"error": err.Error()})
}