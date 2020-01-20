package services

import (
	"customerManagementAppServices/interfaces"
	"customerManagementAppServices/models"
)

func GetLastCustomerIdInSearch(dbHandler interfaces.IDBHandler) models.GetLastCustomerIdInSearchModel {
	return func(searchKeyword string) (int, error) {
		query, queryErr := dbHandler.Query(
			"SELECT customer_id FROM customers WHERE customer_name LIKE ? ORDER BY customer_id DESC",
			"%"+searchKeyword+"%",
		)
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
