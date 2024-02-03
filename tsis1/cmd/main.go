package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"tsis1/internal"
)

func main() {
	fmt.Println("Starting the server...")

	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/anime", handlers.GetAnimeList).Methods("GET")
	r.HandleFunc("/anime/{id:[0-9]+}", handlers.GetAnimeDetails).Methods("GET")
	r.HandleFunc("/health", handlers.HealthCheck).Methods("GET")
  	fmt.Println("Routes registered. Starting the server...")
	http.Handle("/", r)

	if err := http.ListenAndServe(":8081", nil); err != nil {
        	fmt.Println("Error starting the server:", err)
    }
}
