package services

import (
	"customerManagementAppServices/interfaces"
	"customerManagementAppServices/models"
)

// GetSingleCustomerById return a function which fetch a single customer from database by customer_id
func GetSingleCustomerById(dbHandler interfaces.IDBHandler) models.GetSingleCustomerById {
	return func(customerId int) (*models.CustomerModel, error) {
		query, queryErr := dbHandler.Query(
			`
      SELECT customers.customer_id , customer_name , customer_information
      FROM customers , customers_information
      WHERE customers.customer_id = customers_information.customer_id AND
      customers.customer_id = ?
      `, customerId,
		)
		if queryErr != nil {
			return nil, queryErr
		}
		defer query.Close()

		var customer *models.CustomerModel
		if query.Next() {
			customer = &models.CustomerModel{}
			query.Scan(&customer.Id, &customer.Name, &customer.Information)
		}

		return customer, nil
	}
}
