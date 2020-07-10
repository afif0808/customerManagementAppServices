package services

import (
	"customerManagementAppServices/interfaces"
	"customerManagementAppServices/models"
	"database/sql"
)

// SearchCustomers return function which fetch customers with name contains given keyword
func SearchCustomers(dbHandler interfaces.IDBHandler) models.SearchCustomersModel {
	return func(limit int, offset int, keyword string) ([]models.CustomerModel, error) {
		var customers []models.CustomerModel
		query, queryErr := dbHandler.Query(
			`SELECT customer_id , customer_name , customer_information ,customer_addedat
			 FROM customer
       WHERE
			 customer_name LIKE ?
			 ORDER BY customer_id DESC
			 LIMIT ? OFFSET ? `, "%"+keyword+"%", limit, offset)
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
