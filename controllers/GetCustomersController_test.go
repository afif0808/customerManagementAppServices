package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetCustomerController(t *testing.T) {
	testRecoder := httptest.NewRecorder()
	testRouter := mux.NewRouter()
	testRequest := new(http.Request)
}
