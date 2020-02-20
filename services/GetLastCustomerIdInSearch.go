package services

import (
	"customerManagementAppServices/interfaces"
	"customerManagementAppServices/models"
	"database/sql"
)

// GetLastCustomerIdInSearch return a function which fetchs last customer's id
// in descending order which its name contains given keyword from database
func GetLastCustomerIdInSearch(dbHandler interfaces.IDBHandler) models.GetLastCustomerIdInSearchModel {
	return func(searchKeyword string) (int, error) {
		query, queryErr := dbHandler.Query(
			"SELECT customer_id FROM customers WHERE customer_name LIKE ? ORDER BY customer_id DESC",
			"%"+searchKeyword+"%",
		)
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
