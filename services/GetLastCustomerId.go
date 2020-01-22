package services

import (
	"customerManagementAppServices/interfaces"
	"customerManagementAppServices/models"
	"database/sql"
)

// GetLastCustomerId  return function which fetchs last inserted customer from database
func GetLastCustomerId(dbHandler interfaces.IDBHandler) models.GetLastCustomerIdModel {
	return func() (int, error) {
		query, queryErr := dbHandler.Query("SELECT customer_id FROM customers ORDER BY customer_id DESC")
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
