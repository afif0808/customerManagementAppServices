package services

import (
	"customerManagementAppServices/interfaces"
	"customerManagementAppServices/models"
)

func GetSingleCustomer(dbHandler interfaces.IDBHandler) models.GetSingleCustomerById {
	return func(customerId int) (models.CustomerModel, error) {
		query, queryErr := dbHandler.Query(
			`
      SELECT * FROM customers , customers_information
      WHERE
      customers.customer_id = customers_information.customer_id AND
      customer_id = ?
      `, customerId,
		)
		if queryErr != nil {
			return (models.CustomerModel{}), queryErr
		}
		customer := models.CustomerModel{}
		if query.Next() {
			query.Scan(&customer.Id, &customer.Name, &customer.Information)
		}
		return customer, nil
	}
}
