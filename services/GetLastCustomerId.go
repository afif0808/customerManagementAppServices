package services

import (
	"customerManagementAppServices/interfaces"
	"customerManagementAppServices/models"
	"database/sql"
)

// GetLastCustomerId  return a function which fetchs last customer's id in descending order from database
func GetLastCustomerId(dbHandler interfaces.IDBHandler) models.GetLastCustomerIdModel {
	return func() (int, error) {
		query, queryErr := dbHandler.Query("SELECT customer_id FROM customers ORDER BY customer_id DESC")
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
