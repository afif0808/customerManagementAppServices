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

func getMysqlConn() *sql.DB {
	mysqlConn, sqlOpenErr := sql.Open("mysql", "root:@tcp(localhost:3306)/cusandship")
	mysqlConn.SetMaxOpenConns(500)
	if sqlOpenErr != nil {
		log.Fatal(sqlOpenErr)
	}
	return mysqlConn
}
func InitGetCustomersController() http.Handler {
	mysqlHandler := infrastructures.MysqlHandler{Conn: getMysqlConn()}
	getCustomers := services.GetCustomers(&mysqlHandler)
	searchCustomers := services.SearchCustomers(&mysqlHandler)
	getFirstCustomerId := services.GetFirstCustomerId(&mysqlHandler)
	getFirstCustomerIdInSearch := services.GetFirstCustomerIdInSearch(&mysqlHandler)
	return controllers.GetCustomersController(
		getCustomers, searchCustomers, getFirstCustomerId, getFirstCustomerIdInSearch,
	)
}
func InitGetSingleCustomerByIdController() http.Handler {
	mysqlHandler := infrastructures.MysqlHandler{Conn: getMysqlConn()}
	getSingleCustomerById := services.GetSingleCustomerById(&mysqlHandler)
	return controllers.GetSingleCustomerByIdController(getSingleCustomerById)
}
func InitAddCustomerController() http.Handler {
	mysqlHandler := infrastructures.MysqlHandler{Conn: getMysqlConn()}
	addCustomer := services.AddCustomer(&mysqlHandler)
	return controllers.AddCustomerController(addCustomer)
}

func InitUpdateCustomerController() http.Handler {

	mysqlHandler := infrastructures.MysqlHandler{Conn: getMysqlConn()}
	updateCustomer := services.UpdateCustomer(&mysqlHandler)
	return controllers.UpdateCustomerController(updateCustomer)
}
func InitDeleteCustomerController() http.Handler {
	mysqlHandler := infrastructures.MysqlHandler{Conn: getMysqlConn()}
	deleteCustomer := services.DeleteCustomer(&mysqlHandler)
	return cors.AllowAll().Handler(controllers.DeleteCustomerController(deleteCustomer))
}
