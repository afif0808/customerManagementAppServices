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
func GetCustomersController(
	getCustomers models.GetCustomersServiceModel,
	searchCustomers models.SearchCustomersServiceModel,
	getLastCustomerId models.GetLastCustomerIdModel,
	getLastCustomerIdInSearch models.GetLastCustomerIdInSearchModel,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// denoting limit count of customer
		customerLimitCount, parseIntErr := strconv.Atoi(r.URL.Query().Get("limit"))
		if parseIntErr != nil {
			customerLimitCount = 10 // default value
		}

		// if 'offset' parameter is not given by client or it's invalid then it remains '0'
		offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

		searchQuery := r.URL.Query().Get("s")

		var result []models.CustomerModel
		var lastCustomerId int
		var searchCustomersError error
		var getCustomersError error
		var getLastCustomerIdError error
		var getLastCustomerIdInSearchError error

		urlEscapePath := r.URL.EscapedPath() + "?"

		// denoting if client do customer search or get list of customer
		if searchQuery == "" {

			result, getCustomersError = getCustomers(customerLimitCount, offset)
			if getCustomersError != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			// get last customer id to determine the end of page.

			lastCustomerId, getLastCustomerIdError = getLastCustomerId()
			if getLastCustomerIdError != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}

		} else {
			result, searchCustomersError = searchCustomers(customerLimitCount, offset, searchQuery)
			if searchCustomersError != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}

			lastCustomerId, getLastCustomerIdInSearchError = getLastCustomerIdInSearch(searchQuery)
			if getLastCustomerIdInSearchError != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			urlEscapePath += "s=" + searchQuery
		}

		var nextPageLink string
		var previousPageLink string

		if offset > 0 {
			previousPageLink = fmt.Sprintf("%vlimit=%v&offset=%v", urlEscapePath, customerLimitCount, offset-customerLimitCount)
		}
		if len(result) > 0 && lastCustomerId > result[len(result)-1].Id {
			nextPageLink = fmt.Sprintf("%vlimit=%v&offset=%v", urlEscapePath, customerLimitCount, offset+customerLimitCount)
		}

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
