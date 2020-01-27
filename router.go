package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// ServeRoutes initates a router and serves routes
func ServeRoutes() *mux.Router {

	router := mux.NewRouter()
	router.Handle("/api/customers", InitGetCustomersController()).Methods("GET")
	router.Handle("/api/customers/{id}", InitGetSingleCustomerByIdController()).Methods("GET")
	router.Handle("/api/customers", InitAddCustomerController()).Methods("POST")
	return router
}
