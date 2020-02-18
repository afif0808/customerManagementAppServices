package controllers

import (
	"customerManagementAppServices/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

// GetCustomersController gets or searches customers.
// search customer if 's' url parameter is given otherwise it gets customers.
// it implements offset pagination
func GetCustomersController(
	getCustomers models.GetCustomersModel,
	searchCustomers models.SearchCustomersModel,
	getFirstCustomerId models.GetFirstCustomerIdModel,
	getFirstCustomerIdInSearch models.GetFirstCustomerIdInSearchModel,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		// denoting limit count of customer
		customerLimitCount, parseIntErr := strconv.Atoi(r.URL.Query().Get("limit"))
		if parseIntErr != nil {
			customerLimitCount = 10 // default value
		}
		if customerLimitCount < 1 {
			w.WriteHeader(http.StatusBadRequest)
		}

		// denoting customer list offset
		// if 'offset' parameter is not given by client or it's invalid then it remains '0'
		offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

		if offset < 0 {
			w.WriteHeader(http.StatusBadRequest)
		}

		searchQuery := r.URL.Query().Get("s")

		var result []models.CustomerModel
		var customerListLowerBound int
		var searchCustomersError error
		var getCustomersError error
		var getFirstCustomerIdError error
		var getFirstCustomerIdInSearchError error

		siteURL := fmt.Sprintf("http://%v%v?", r.Host, r.URL.Path)
		// denoting if client do customer search or only get list of customer
		if searchQuery == "" {
			// get customers from database
			result, getCustomersError = getCustomers(customerLimitCount, offset)
			if getCustomersError != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// get first customer id as a customer list lower bound
			// to check if end of page is reached
			customerListLowerBound, getFirstCustomerIdError = getFirstCustomerId()
			if getFirstCustomerIdError != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

		} else {
			// search customers from database
			result, searchCustomersError = searchCustomers(customerLimitCount, offset, searchQuery)
			if searchCustomersError != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// get first customer id as a customer list lower bound
			// to check if end of page is reached
			customerListLowerBound, getFirstCustomerIdInSearchError = getFirstCustomerIdInSearch(searchQuery)
			if getFirstCustomerIdInSearchError != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			siteURL += "s=" + searchQuery + "&"
		}

		var nextPageLink string
		var previousPageLink string

		// denoting next and previous customer list link
		if offset > 0 {
			previousPageLink = fmt.Sprintf("%vlimit=%v&offset=%v", siteURL, customerLimitCount, offset-customerLimitCount)
		}

		if len(result) > 0 && customerListLowerBound < result[len(result)-1].Id {
			nextPageLink = fmt.Sprintf("%vlimit=%v&offset=%v", siteURL, customerLimitCount, offset+customerLimitCount)
		}

		APIResponse := models.BulkCustomersAPIModel{
			NextPageLink:     nextPageLink,
			PreviousPageLink: previousPageLink,
			Result:           result,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(APIResponse)
	})
}
