package controllers

import (
	"customerManagementAppServices/models"
	"encoding/json"
	"fmt"
	"net/http"
)

// AddCustomerController adds new customer
// controller of AddCustomer service .
func AddCustomerController(addCustomer models.AddCustomerModel) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		postData := map[string]interface{}{}

		json.NewDecoder(r.Body).Decode(&postData)

		customerName := fmt.Sprint(postData["customerName"])
		customerInformation := fmt.Sprint(postData["customerInformation"])

		// fmt.Fprintln(os.Stdout, customerName, customerInformation)
		if customerName == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		addCustomerError := addCustomer(customerName, customerInformation)
		if addCustomerError != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}
