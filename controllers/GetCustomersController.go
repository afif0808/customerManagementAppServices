package controllers

import (
	"customerManagementAppServices/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// GetCustomersController search for customers by name if parameter url parameter 'q' is defined
// otherwise , it searves list of customers
// implements 'Seek Pagination' method
func GetCustomersController(getCustomers models.GetCustomersServiceModel, searchCustomers models.SearchCustomersServiceModel) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// denoting limit count of customer
		customerLimitCount, parseIntErr := strconv.Atoi(r.URL.Query().Get("limit"))
		if parseIntErr != nil {
			customerLimitCount = 10 // default value
		}

		// if 'offset' parameter is not given by client or it's invalid then it remains '0'
		offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

		searchQuery := r.URL.Query().Get("q")

		var result []models.CustomerModel
		var searchCustomersError error
		var getCustomersError error

		// denoting if client do customer search or simply get list of customer

		if searchQuery == "" {
			result, getCustomersError = getCustomers(customerLimitCount, offset)
			if getCustomersError != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		} else {
			result, searchCustomersError = searchCustomers(customerLimitCount, offset, searchQuery)
			if searchCustomersError != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}

		var nextPageLink string
		var previousPageLink string

		if offset > 1 && len(result) > 0 {
			previousPageLink = fmt.Sprintf("%v?limit=%v&offset=%v", r.URL.String(), customerLimitCount, offset-customerLimitCount)
		}
		nextPageLink = fmt.Sprintf("%v?limit=%v&offset=%v", r.URL.String(), customerLimitCount, result[len(result)-1].Id)

		APIResponse := models.BulkCustomersAPIModel{
			CustomerLimitCount: customerLimitCount,
			NextPageLink:       nextPageLink,
			PreviousPageLink:   previousPageLink,
			Result:             result,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(APIResponse)
	})
}
