package services

import (
	"customerManagementAppServices/interfaces"
	"customerManagementAppServices/models"
	"database/sql"
)

// GetLastCustomerIdInSearch return a function which fetchs first customer's id
// in ascending order which its name contains given keyword from database
func GetFirstCustomerIdInSearch(dbHandler interfaces.IDBHandler) models.GetFirstCustomerIdInSearchModel {
	return func(searchKeyword string) (int, error) {
		query, queryErr := dbHandler.Query(
			"SELECT customer_id FROM customers WHERE customer_name LIKE ? ORDER BY customer_id ASC",
			"%"+searchKeyword+"%",
		)
		if queryErr != nil && queryErr != sql.ErrNoRows {
			return 0, queryErr
		}
		var customerId int
		if query.Next() {
			query.Scan(&customerId)
		}

		return customerId, nil
	}
}
