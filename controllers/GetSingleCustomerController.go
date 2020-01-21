package controllers

import (
	"customerManagementAppServices/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetSingleCustomerByIdController serve a single customer
func GetSingleCustomerByIdController(getSingleCustomerById models.GetSingleCustomerById) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		customerId, parseIntError := strconv.Atoi(mux.Vars(r)["id"])
		if parseIntError != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		customer, getSingleCustomerByIdError := getSingleCustomerById(customerId)
		if getSingleCustomerByIdError != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		apiResponse := models.SingleCustomerAPIModel{
			Result: customer,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(apiResponse)
	})
}
