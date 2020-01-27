package services

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestAddCustomer(t *testing.T) {
	// mysqlConn, sqlOpenErr := sql.Open("mysql", "root:@tcp(localhost:3306)/customer_management")
	// if sqlOpenErr != nil {
	// 	log.Fatal(sqlOpenErr)
	// }
	// mysqlHandler := infrastructures.MysqlHandler{Conn: mysqlConn}
	// addCustomer := AddCustomer(&mysqlHandler)
	// // // addCustomerErr := addCustomer("Imen", "Tukang Angkek Kosiak")
	// // log.Println(addCustomerErr)
}
