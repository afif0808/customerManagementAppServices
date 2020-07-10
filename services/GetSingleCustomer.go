package services

import (
	"customerManagementAppServices/interfaces"
	"customerManagementAppServices/models"
)

// GetSingleCustomerById return a function which fetch a single customer from database by customer_id
func GetSingleCustomerById(dbHandler interfaces.IDBHandler) models.GetSingleCustomerById {
	return func(customerId int) (*models.CustomerModel, error) {
		query, queryErr := dbHandler.Query(
			`SELECT customer_id , customer_name , customer_information ,customer_addedat
			 FROM customer
			 ORDER BY customer_id DESC
      customer.customer_id = ?
      `, customerId,
		)
		if queryErr != nil {
			return nil, queryErr
		}
		defer query.Close()

		var customer *models.CustomerModel
		if query.Next() {
			customer = &models.CustomerModel{}
			query.Scan(&customer.Id, &customer.Name, &customer.Information, &customer.DateAdded)
		}

		return customer, nil
	}
}
