package services

import (
	"customerManagementAppServices/interfaces"
	"customerManagementAppServices/models"
	"database/sql"
)

// GetLastCustomerIdInSearch return function which fetchs last inserted customer
// where name contains given keyword from database
func GetLastCustomerIdInSearch(dbHandler interfaces.IDBHandler) models.GetLastCustomerIdInSearchModel {
	return func(searchKeyword string) (int, error) {
		query, queryErr := dbHandler.Query(
			"SELECT customer_id FROM customers WHERE customer_name LIKE ? ORDER BY customer_id DESC",
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
