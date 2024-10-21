package main

import "net/http"

func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithErro(w, 400, "someting went worng")
}
