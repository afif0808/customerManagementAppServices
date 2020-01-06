package controllers

import (
	"customerManagementAppServices/models"
	"net/http"
	"strconv"
)

// GetCustomersController search for customers by name if parameter url parameter 'q' is defined
// otherwise , it serves list of customers
// implements 'Seek Pagination' method.
func GetCustomersController(getCustomers models.GetCustomersFunc, searchCustomers models.SearchCustomersFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// get 'limit' url parameter value and parse it to be an integer.
		// if error is found then client 'limit' probably not defined by client
		// or the type is not valid  , so default value is assigned instead.
		customerLimit, parseIntErr := strconv.Atoi(r.URL.Query().Get("limit"))
		if parseIntErr != nil {
			customerLimit = 10 // '10' is default value
		}

		// get 'afterId' url parameter value and parse it to be an integer
		// if 'afterId' is not given or it's not string number , the  integer conversion will return an error
		// and afterId variable remains zero since it's the default value of integer
		afterId, _ := strconv.Atoi(r.URL.Query().Get("afterId"))

		searchQuery := r.URL.Query().Get("q")

		var result []models.CustomerModel
		var searchCustomersError error
		var getCustomersError error
		if searchQuery == "" {
			result, getCustomersError = getCustomers(customerLimit, afterId)
		} else {
			result, searchCustomersError = searchCustomers(customerLimit, afterId, searchQuery)
		}
	})
}
