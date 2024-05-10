package utils

import (
	"log"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, msg string) {

	if code > 499 {
		log.Printf("Responding with 5XX error: %v", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}
	RespondWithJson(w, code, errResponse{
		Error: msg,
	})

}
