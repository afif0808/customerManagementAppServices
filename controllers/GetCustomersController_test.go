package controllers

import (
	"customerManagementAppServices/models"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetCustomerController(t *testing.T) {
	testRecoder := httptest.NewRecorder()
	testRouter := mux.NewRouter()
	testRouter.Handle("/customers", GetCustomersController(models.GetCustomersFunc{}, models.SearchCustomersFunc))
	testGetCustomers := httptest.NewRequest("GET", "http://localhost:378/customers", nil)
	testSearchCustomers := httptest.NewRequest("GET", "http://localhost:378/customers?q=fulan", nil)
	testPaging := httptest.NewRequest("GET", "http://localhost:378/customers?limit=10,afterId=5", nil)

}
