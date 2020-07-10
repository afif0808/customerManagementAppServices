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
			`SELECT customer_id , customer_name , customer_information ,customer_addedat
			 FROM customer
			 ORDER BY customer_id DESC
			 LIMIT ? OFFSET ?`, limit, offset)

		if queryErr != nil && queryErr != sql.ErrNoRows {
			return nil, queryErr
		}

		defer query.Close()

		for query.Next() {
			customer := models.CustomerModel{}
			query.Scan(&customer.Id, &customer.Name, &customer.Information, &customer.DateAdded)
			customers = append(customers, customer)
		}
		return customers, nil
	}

}
