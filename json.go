package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithErro(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5xx error")
	}
	type erroResponse struct {
		error string `json:"error"`
	}
	responseWithJSON(w, code, erroResponse{
		error: msg,
	})
}

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("fail to marshal json response %v", payload)
		w.WriteHeader(code)
		return
	}
	w.Header().Add("content-type", "application.json")
	w.WriteHeader(code)
	w.Write(data)
}
