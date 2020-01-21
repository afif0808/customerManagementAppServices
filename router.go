package main

import (
	"customerManagementAppServices/controllers"
	"customerManagementAppServices/infrastructures"
	"customerManagementAppServices/services"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// ServeRoutes initates a router and serves routes
func ServeRoutes() *mux.Router {

	router := mux.NewRouter()

	mysqlConn, sqlOpenErr := sql.Open("mysql", "root:@tcp(localhost:3306)/customer_management")

	if sqlOpenErr != nil {
		log.Fatal(sqlOpenErr)
	}

	mysqlHandler := infrastructures.MysqlHandler{Conn: mysqlConn}
	getCustomersService := services.GetCustomerService(&mysqlHandler)
	searchCustomerService := services.SearchCustomerService(&mysqlHandler)
	getLastCustomerId := services.GetLastCustomerId(&mysqlHandler)
	getLastCustomerIdInSearch := services.GetLastCustomerIdInSearch(&mysqlHandler)
	getCustomersController := controllers.GetCustomersController(
		getCustomersService, searchCustomerService, getLastCustomerId, getLastCustomerIdInSearch,
	)
	//routes
	router.Handle("/api/customers", getCustomersController)

	return router
}
