package main

import (
	"dowser/api/handlers"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()
	handler := corsSetup(mux)
	handlers.GetGraph(mux)

	log.Println("Starting API Server on port 8080...")
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func corsSetup(h http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
	})
	return c.Handler(h)
}
