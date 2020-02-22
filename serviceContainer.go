package main

import (
	"customerManagementAppServices/controllers"
	"customerManagementAppServices/infrastructures"
	"customerManagementAppServices/services"
	"database/sql"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func InitGetCustomersController() http.Handler {
	mysqlConn, sqlOpenErr := sql.Open("mysql", "root:@tcp(localhost:3306)/customer_management")
	mysqlConn.SetMaxOpenConns(500)
	if sqlOpenErr != nil {
		log.Fatal(sqlOpenErr)
	}
	mysqlHandler := infrastructures.MysqlHandler{Conn: mysqlConn}
	getCustomers := services.GetCustomers(&mysqlHandler)
	searchCustomers := services.SearchCustomers(&mysqlHandler)
	getFirstCustomerId := services.GetFirstCustomerId(&mysqlHandler)
	getFirstCustomerIdInSearch := services.GetFirstCustomerIdInSearch(&mysqlHandler)
	return controllers.GetCustomersController(
		getCustomers, searchCustomers, getFirstCustomerId, getFirstCustomerIdInSearch,
	)
}
func InitGetSingleCustomerByIdController() http.Handler {
	mysqlConn, sqlOpenErr := sql.Open("mysql", "root:@tcp(localhost:3306)/customer_management")
	mysqlConn.SetMaxOpenConns(500)
	if sqlOpenErr != nil {
		log.Fatal(sqlOpenErr)
	}
	mysqlHandler := infrastructures.MysqlHandler{Conn: mysqlConn}
	getSingleCustomerById := services.GetSingleCustomerById(&mysqlHandler)
	return controllers.GetSingleCustomerByIdController(getSingleCustomerById)
}
func InitAddCustomerController() http.Handler {
	mysqlConn, sqlOpenErr := sql.Open("mysql", "root:@tcp(localhost:3306)/customer_management")
	mysqlConn.SetMaxOpenConns(500)
	if sqlOpenErr != nil {
		log.Fatal(sqlOpenErr)
	}
	mysqlHandler := infrastructures.MysqlHandler{Conn: mysqlConn}
	addCustomer := services.AddCustomer(&mysqlHandler)
	return controllers.AddCustomerController(addCustomer)
}

func InitUpdateCustomerController() http.Handler {
	mysqlConn, sqlOpenErr := sql.Open("mysql", "root:@tcp(localhost:3306)/customer_management")
	mysqlConn.SetMaxOpenConns(500)
	if sqlOpenErr != nil {
		log.Fatal(sqlOpenErr)
	}
	mysqlHandler := infrastructures.MysqlHandler{Conn: mysqlConn}
	updateCustomer := services.UpdateCustomer(&mysqlHandler)
	return controllers.UpdateCustomerController(updateCustomer)
}
func InitDeleteCustomerController() http.Handler {
	mysqlConn, sqlOpenErr := sql.Open("mysql", "root:@tcp(localhost:3306)/customer_management")
	mysqlConn.SetMaxOpenConns(500)

	if sqlOpenErr != nil {
		log.Fatal(sqlOpenErr)
	}
	mysqlHandler := infrastructures.MysqlHandler{Conn: mysqlConn}
	deleteCustomer := services.DeleteCustomer(&mysqlHandler)
	return cors.AllowAll().Handler(controllers.DeleteCustomerController(deleteCustomer))
}
