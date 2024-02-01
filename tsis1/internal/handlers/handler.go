package handlers

import (
	"fmt"
	"net/http"
	"github.com/Altynbek38/Go-2024/tsis1/api"
)

func getAll(w http.ResponseWriter, r *http.Request) {
	drivers := api.Drivers

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(drivers)	
}

func getByFirstName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	driverName = vars["firstname"]

	driver := api.getDriverByFirstName(driverName)
	
	w.Header().Set("Content-Type", "application/json")
	
	if driver == nil {
		w.WriteHeader(http.StatusNotFound())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(driver)
}

func healthCheck (w. http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello to F1 Drivers API\n This API provides data about 20 F1 drivers \n Author: Altynbek Zholdybay")
}