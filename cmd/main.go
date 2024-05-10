package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Shreyank031/go-rss_scraper/pkg/handlers"
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
	v1Router.Get("/foo", handlers.HandlerReadiness)
	v1Router.Get("/error", handlers.HandlerError)
	router.Mount("/v1", v1Router)

	log.Printf("The server is starting at: %v", postString)
	log.Fatal(http.ListenAndServe("localhost:8081", router))
}
