package controllers

import "net/http"
import "customerManagementAppServices/models"

// GetCustomersController search for customers by name if parameter url parameter 'q' is defined
// otherwise , it serves list of customers
func GetCustomersController(getCustomer models.GetCustomersFunc, searchCustomers models.SearchCustomersFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
