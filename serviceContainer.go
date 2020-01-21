package main

import (
	"customerManagementAppServices/controllers"
	"customerManagementAppServices/infrastructures"
	"customerManagementAppServices/services"
	"database/sql"
	"log"
	"net/http"
)

func InitGetCustomersController() http.Handler {
	mysqlConn, sqlOpenErr := sql.Open("mysql", "root:@tcp(localhost:3306)/customer_management")
	if sqlOpenErr != nil {
		log.Fatal(sqlOpenErr)
	}
	mysqlHandler := infrastructures.MysqlHandler{Conn: mysqlConn}
	getCustomers := services.GetCustomers(&mysqlHandler)
	searchCustomers := services.SearchCustomers(&mysqlHandler)
	getLastCustomerId := services.GetLastCustomerId(&mysqlHandler)
	getLastCustomerIdInSearch := services.GetLastCustomerIdInSearch(&mysqlHandler)
	return controllers.GetCustomersController(
		getCustomers, searchCustomers, getLastCustomerId, getLastCustomerIdInSearch,
	)
}
func InitGetSingleCustomerByIdController() http.Handler {
	mysqlConn, sqlOpenErr := sql.Open("mysql", "root:@tcp(localhost:3306)/customer_management")
	if sqlOpenErr != nil {
		log.Fatal(sqlOpenErr)
	}
	mysqlHandler := infrastructures.MysqlHandler{Conn: mysqlConn}
	getSingleCustomerById := services.GetSingleCustomerById(&mysqlHandler)
	return controllers.GetSingleCustomerByIdController(getSingleCustomerById)
}
