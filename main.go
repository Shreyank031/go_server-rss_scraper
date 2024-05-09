package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error while reading env file: %v", err)
	}

	postString := os.Getenv("PORT")
	fmt.Println("Hello, Golang!")
	if postString != "" {
		fmt.Println("Port: ", postString)
	} else {
		fmt.Println("Port not available")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUt", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/foo", handlerReadiness)
	v1Router.Get("/error", handlerError)
	router.Mount("/v1", v1Router)

	log.Printf("The server is starting at: %v", postString)
	log.Fatal(http.ListenAndServe("localhost:8081", router))
}

func handlerReadiness(w http.ResponseWriter, h *http.Request) {
	respondWithJson(w, http.StatusOK, struct{}{})

}
func handlerError(w http.ResponseWriter, h *http.Request) {

	respondWithError(w, 400, "Something went wrong!")

}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {

	dat, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Failed to marshal JSON response: %v", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code) //same  as 200
	w.Write(dat)

}

func respondWithError(w http.ResponseWriter, code int, msg string) {

	if code > 499 {
		log.Printf("Responding with 5XX error: %v", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}
	respondWithJson(w, code, errResponse{
		Error: msg,
	})

}
