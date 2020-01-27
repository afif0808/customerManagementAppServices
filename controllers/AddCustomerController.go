package controllers

import (
	"customerManagementAppServices/models"
	"net/http"
)

// AddCustomerController adds new customer to database
// controller of AddCustomer service .
func AddCustomerController(addCustomer models.AddCustomerModel) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		customerName := r.PostFormValue("customerName")
		customerInformation := r.PostFormValue("customerInformation")
		if customerName == "" {
			w.WriteHeader(http.StatusBadRequest)
		}
		addCustomerError := addCustomer(customerName, customerInformation)
		if addCustomerError != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
	})
}
