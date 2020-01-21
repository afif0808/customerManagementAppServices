package controllers

import (
	"customerManagementAppServices/models"
	"encoding/json"
	"net/http"
	"strconv"
)

// GetSingleCustomerByIdController serve a single customer by id
func GetSingleCustomerByIdController(getSingleCustomerById models.GetSingleCustomerById) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		customerId, parseIntError := strconv.Atoi(r.URL.Query().Get("id"))
		if parseIntError != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		customer, getSingleCustomerByIdError := getSingleCustomerById(customerId)
		if getSingleCustomerByIdError != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		apiResponse := models.SingleCustomerAPIModel{
			Result: customer,
		}
		json.NewEncoder(w).Encode(apiResponse)
	})
}
