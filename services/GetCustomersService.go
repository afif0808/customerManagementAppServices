package services

import (
	"customerManagementAppServices/interfaces"
	"customerManagementAppServices/models"
)

// GetCustomersService returns function which fetchs (from database) to at most 'limit' customers
// starting from customer number 'offset'
// the customers is sorted by id

func GetCustomerService(dbHandler interfaces.IDBHandler) models.GetCustomersServiceModel {
	return func(limit int, offset int) ([]models.CustomerModel, error) {
		var customers []models.CustomerModel
		query, queryErr := dbHandler.Query(
			`SELECT customers.customer_id , customer_name , customer_information FROM customers
			 INNER JOIN customers_information ON customers.customer_id = customers_information.customer_id
			 LIMIT ? OFFSET ?`, limit, offset)

		if queryErr != nil {
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
