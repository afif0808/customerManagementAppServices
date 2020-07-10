package services

import (
	"customerManagementAppServices/interfaces"
	"customerManagementAppServices/models"
)

// DeleteCustomer delete customer from database
// 'customerId' parameter is in string type for simplicity
// because anyway it works.
func DeleteCustomer(dbHandler interfaces.IDBHandler) models.DeleteCustomerModel {
	return func(customerId string) error {
		deleteCustomerError := dbHandler.Execute("DELETE FROM customer WHERE customer_id = ?", customerId)
		if deleteCustomerError != nil {
			return deleteCustomerError
		}
		return nil
	}
}
