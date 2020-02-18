package services

import (
	"customerManagementAppServices/interfaces"
	"customerManagementAppServices/models"
	"database/sql"
)

// GetCustomersService gets customers from database . the customers are sorted by id.
// this function implements 'Offset Pagination'
// it's done by returning function that implement GetCustomersModel.
func GetCustomers(dbHandler interfaces.IDBHandler) models.GetCustomersModel {
	return func(limit int, offset int) ([]models.CustomerModel, error) {
		var customers []models.CustomerModel
		query, queryErr := dbHandler.Query(
			`SELECT customers.customer_id , customer_name , customer_information
			 FROM customers , customers_information
			 WHERE customers.customer_id = customers_information.customer_id
			 ORDER BY customers.customer_id DESC
			 LIMIT ? OFFSET ?`, limit, offset)

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
