package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5XX error:", msg)
	}
	type errResponce struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errResponce{Error: msg})
}
func respondWithJSON(w http.ResponseWriter, code int, playload interface{}) {
	dat, err := json.Marshal(playload)
	if err != nil {
		log.Printf("Failed to marshal JSON responce: %v", playload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
