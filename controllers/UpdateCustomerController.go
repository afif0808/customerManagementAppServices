package controllers

import (
	"customerManagementAppServices/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// UpdateCustomerController update customer name and information
func UpdateCustomerController(updateCustomer models.UpdateCustomerModel) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		postData := map[string]interface{}{}

		// reading posted data
		json.NewDecoder(r.Body).Decode(&postData)

		// for simplicity customerId is in string type because any type works fine in this case
		customerId := mux.Vars(r)["id"]

		newCustomerName := fmt.Sprint(postData["newCustomerName"])
		newCustomerInformation := fmt.Sprint(postData["newCustomerInformation"])
		log.Println("nah loh ", newCustomerName, newCustomerInformation)
		if customerId == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		updateCustomerError := updateCustomer(customerId, newCustomerName, newCustomerInformation)
		if updateCustomerError != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}
