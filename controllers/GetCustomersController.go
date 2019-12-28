package controllers

import "net/http"
import "customerManagementAppServices/models"

func GetCustomersController(getCustomer models.GetCustomersFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
