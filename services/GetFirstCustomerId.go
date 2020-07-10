package services

import (
	"customerManagementAppServices/interfaces"
	"customerManagementAppServices/models"
	"database/sql"
)

// GetLastCustomerId  return a function which fetchs first customer's id in ascending order from database
func GetFirstCustomerId(dbHandler interfaces.IDBHandler) models.GetFirstCustomerIdModel {
	return func() (int, error) {
		query, queryErr := dbHandler.Query("SELECT customer_id FROM customer ORDER BY customer_id ASC")
		if queryErr != nil && queryErr != sql.ErrNoRows {
			return 0, queryErr
		}
		defer query.Close()
		var customerId int
		if query.Next() {
			query.Scan(&customerId)
		}

		return customerId, nil
	}
}
