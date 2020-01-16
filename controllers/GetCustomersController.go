package controllers

import (
	"customerManagementAppServices/models"
	"encoding/json"
	"net/http"
	"strconv"
)

// GetCustomersController search for customers by name if parameter url parameter 'q' is defined
// otherwise , it serves list of customers
// implements 'Seek Pagination' method
func GetCustomersController(getCustomers models.GetCustomersServiceModel, searchCustomers models.SearchCustomersServiceModel) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// denoting limit count of customer
		customerLimit, parseIntErr := strconv.Atoi(r.URL.Query().Get("limit"))
		if parseIntErr != nil {
			customerLimit = 10 // default value
		}

		// if 'afterId' parameter is not given by client or it's invalid then it remains '0'
		afterId, _ := strconv.Atoi(r.URL.Query().Get("afterId"))

		searchQuery := r.URL.Query().Get("q")

		var result []models.CustomerModel
		var searchCustomersError error
		var getCustomersError error
		if searchQuery == "" {
			result, getCustomersError = getCustomers(customerLimit, afterId)
			if getCustomersError != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		} else {
			result, searchCustomersError = searchCustomers(customerLimit, afterId, searchQuery)
			if searchCustomersError != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}

		json.NewEncoder(w).Encode(result)
	})
}
