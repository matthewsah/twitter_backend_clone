package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respond_with_error(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5xx error:", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	respond_with_JSON(w, code, errResponse{
		Error: msg,
	})
}

func respond_with_JSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
