package main

import (
	"encoding/json"
	"log"
	"net/http"
)

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
