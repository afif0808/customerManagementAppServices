package services

import (
	"customerManagementAppServices/interfaces"
	"customerManagementAppServices/models"
)

// DeleteCustomer delete customer from database
// 'customerId' parameter is in string type for simplicity because there's no point on converting
// it to integer
func DeleteCustomer(dbHandler interfaces.IDBHandler) models.DeleteCustomerModel {
	return func(customerId string) error {
		deleteCustomerError := dbHandler.Execute("DELETE FROM customers WHERE customer_id = ?", customerId)
		deleteCustomerInformationError := dbHandler.Execute("DELETE FROM customers_information WHERE customer_id = ?", customerId)
		if deleteCustomerError != nil {
			return deleteCustomerError
		}
		if deleteCustomerInformationError != nil {
			return deleteCustomerInformationError
		}
		return nil
	}
}
