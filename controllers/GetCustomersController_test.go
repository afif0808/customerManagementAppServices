package controllers

import (
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetCustomerController(t *testing.T) {
	testRecoder := httptest.NewRecorder()
	testRouter := mux.NewRouter()
	testRequest := httptest.NewRequest("GET", "http://localhost:378/customers", nil)
}
