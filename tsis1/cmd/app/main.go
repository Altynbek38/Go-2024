package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	log.Println("Starting")
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/drivers", GetAll).Methods("GET")
	router.HandleFunc("/drivers/{firstname}", GetByFirstName).Methods("GET")
	log.Println("Processing")
	http.ListenAndServe(":8080", router)

}