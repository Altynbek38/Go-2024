package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	formula "github.com/Altynbek38/Go-2024/tsis1/internal/app"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	drivers := formula.Drivers

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(drivers)	
}

func GetByFirstName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	driverName := vars["firstname"]

	driver := formula.GetDriverByFirstName(driverName)
	
	w.Header().Set("Content-Type", "application/json")
	
	if driver.Firstname == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(driver)
}

func HealthCheck (w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello to F1 Drivers API\n This API provides data about 20 F1 drivers \n Author: Altynbek Zholdybay")
}