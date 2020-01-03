package controllers

import (
	"customerManagementAppServices/models"
	"net/http"
	"strconv"
)

// GetCustomersController search for customers by name if parameter url parameter 'q' is defined
// otherwise , it serves list of customers
func GetCustomersController(getCustomer models.GetCustomersFunc, searchCustomers models.SearchCustomersFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// get 'limit' url query and parse it into integer.
		// if error is found then client 'limit' probably not defined by client
		// or the type is not valid so default value is assigned instead.
		limit, parseIntErr := strconv.Atoi(r.URL.Query().Get("limit"))
		if parseIntErr != nil {
			limit = 10 // '10' is default value
		}

	})
}
