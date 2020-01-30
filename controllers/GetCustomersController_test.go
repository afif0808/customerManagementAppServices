package controllers

import (
	"customerManagementAppServices/infrastructures"
	"customerManagementAppServices/services"
	"database/sql"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetCustomerController(t *testing.T) {
	testRecoder := httptest.NewRecorder()
	testRouter := mux.NewRouter()

	mysqlConn, sqlOpenErr := sql.Open("mysql", "root:@tcp(localhost:3306)/customer_management")
	if sqlOpenErr != nil {
		log.Fatal(sqlOpenErr)
	}
	mysqlHandler := infrastructures.MysqlHandler{Conn: mysqlConn}
	getCustomersService := services.GetCustomers(&mysqlHandler)
	searchCustomerService := services.SearchCustomers(&mysqlHandler)
	getFirstCustomerId := services.GetFirstCustomerId(&mysqlHandler)
	getFirstCustomerIdInSearch := services.GetFirstCustomerIdInSearch(&mysqlHandler)
	getCustomersController := GetCustomersController(
		getCustomersService, searchCustomerService, getFirstCustomerId, getFirstCustomerIdInSearch,
	)
	//routes
	testRouter.Handle("/customers", getCustomersController)

	// testRouter.Handle("/customers", GetCustomersController(models.GetCustomersFunc{}, models.SearchCustomersFunc{}))
	// testGetCustomers := httptest.NewRequest("GET", "http://localhost:378/customers", nil)
	testSearchCustomers := httptest.NewRequest("GET", "http://localhost:378/customers?s=fulan", nil)
	// testPaging := httptest.NetwRequest("GET", "http://localhost:378/customers?limit=10,afterId=5", nil)
	testRouter.ServeHTTP(testRecoder, testSearchCustomers)

}
