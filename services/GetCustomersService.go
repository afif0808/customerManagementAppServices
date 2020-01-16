package services

import (
	"customerManagementAppServices/interfaces"
	"customerManagementAppServices/models"
)

// GetCustomersService returns function which fetchs (from database) to at most 'limit' customers
// starting from 'afterId' customer
// the customers is sorted by id

func GetCustomerService(dbHandler interfaces.IDBHandler) models.GetCustomersServiceModel {
	return func(limit int, afterId int) ([]models.CustomerModel, error) {
		var customers []models.CustomerModel
		query, queryErr := dbHandler.Query("SELECT * FROM customers WHERE customer_id > ? LIMIT ?", limit, afterId)
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
