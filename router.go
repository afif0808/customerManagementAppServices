package main

import (
	"github.com/gorilla/mux"
)

// ServeRoutes initate a router and routes
func ServeRoutes() *mux.Router {
	router := mux.NewRouter()
	//routes
	// router.Handle("/api/customers", )

	return router
}
