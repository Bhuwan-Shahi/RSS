package main

import "net/http"

func hadleReadiness(w http.ResponseWriter, r *http.Request) {
	responseWithJSON(w, 200, struct{}{})
}
