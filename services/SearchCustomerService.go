package services

import (
	"customerManagementAppServices/interfaces"
	"customerManagementAppServices/models"
)

func SearchCustomerService(dbHandler interfaces.IDBHandler) models.SearchCustomersServiceModel {
	return func(limit int, offset int, keyword string) ([]models.CustomerModel, error) {
		var customers []models.CustomerModel
		query, queryErr := dbHandler.Query(
			`SELECT customers.customer_id , customer_name , customer_information FROM customers
       INNER JOIN customers_information ON customers.customer_id = customers_information.customer_id
       WHERE customer_name LIKE ? LIMIT ? OFFSET ?`, "%"+keyword+"%", limit, offset)
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
