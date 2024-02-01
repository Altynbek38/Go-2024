package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/Altynbek38/Go-2024/tsis1/internal/handlers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/health-check", healthCheck).Methods("GET")
	router.HandleFunc("/drivers", getAll).Methods("GET")
	router.HandleFunc("/drivers/{firstname}", getByFirstName).Methods("GET")

	http.ListenAndServe(":80", router)
}