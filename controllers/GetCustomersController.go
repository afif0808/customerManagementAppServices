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

		// see if limit is defined or is '0'
		// var limit will be '0' if it's not defined
		// if limit is not defined then 'strconv.Atoi' will return error and var limit is zero
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
		if limit == 0 {
			// default limit
			limit = 10
		}

	})
}
