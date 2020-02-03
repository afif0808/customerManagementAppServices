package controllers

import (
	"customerManagementAppServices/models"
	"net/http"

	"github.com/gorilla/mux"
)

// DeleteCustomerController deletes a customer
// controller of DeleteCustomer service
func DeleteCustomerController(deleteCustomer models.DeleteCustomerModel) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// for simplicity customerId is in string type because any type works fine in this case
		customerId := mux.Vars(r)["id"]

		deleteCustmerError := deleteCustomer(customerId)
		if deleteCustmerError != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
	})
}
