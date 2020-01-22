package services

import (
	"customerManagementAppServices/interfaces"
	"customerManagementAppServices/models"
	"database/sql"
)

// SearchCustomers return function which fetch customers which
// its name contains given keyword from database
func SearchCustomers(dbHandler interfaces.IDBHandler) models.SearchCustomersModel {
	return func(limit int, offset int, keyword string) ([]models.CustomerModel, error) {
		var customers []models.CustomerModel
		query, queryErr := dbHandler.Query(
			`SELECT customers.customer_id , customer_name , customer_information
			 FROM customers,customers_information
       WHERE
			 customers.customer_id = customers_information.customer_id AND
			 customer_name LIKE ?
			 ORDER BY customers.customer_id DESC
			 LIMIT ? OFFSET ? `, "%"+keyword+"%", limit, offset)
		if queryErr != nil && queryErr != sql.ErrNoRows {
			return nil, queryErr
		}
		for query.Next() {
			customer := models.CustomerModel{}
			query.Scan(&customer.Id, &customer.Name, &customer.Information)
			customers = append(customers, customer)
		}
		return customers, nil
	}
}
