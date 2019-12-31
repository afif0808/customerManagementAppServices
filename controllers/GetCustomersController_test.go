package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)


func TestGetCustomerController(t *testing.T) {
  testRecoder := httptest.NewRecorder()
  testRouter := mux.NewRouter()
  testRequest := new(http.Request)
}
