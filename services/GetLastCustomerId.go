package services

import (
	"customerManagementAppServices/interfaces"
	"customerManagementAppServices/models"
)

func GetLastCustomerId(dbHandler interfaces.IDBHandler) models.GetLastCustomerIdModel {
	return func() (int, error) {
		query, queryErr := dbHandler.Query("SELECT customer_id FROM customers ORDER BY customer_id DESC")
		if queryErr != nil {
			return 0, queryErr
		}
		var customerId int
		query.Next()
		scanErr := query.Scan(&customerId)
		if scanErr != nil {
			return 0, scanErr
		}
		return customerId, nil
	}
}
