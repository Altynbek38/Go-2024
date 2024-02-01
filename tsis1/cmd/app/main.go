package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/Altynbek38/Go-2024/tsis1/internal/handlers"
)

func main() {
	router := mux.NewRouter()
	log.Println("Starting")
	router.HandleFunc("/health-check", handlers.HealthCheck).Methods("GET")
	router.HandleFunc("/drivers", handlers.GetAll).Methods("GET")
	router.HandleFunc("/drivers/{firstname}", handlers.GetByFirstName).Methods("GET")
	log.Println("Processing")
	http.ListenAndServe(":8080", router)

}