package services

import (
	"customerManagementAppServices/interfaces"
	"customerManagementAppServices/models"
	"fmt"
)

// UpdateCustomer return function which updates customer's name and information in database
func UpdateCustomer(dbHandler interfaces.IDBHandler) models.UpdateCustomerModel {
	return func(customerId, newCustomerName, newCustomerInformation string) error {
		if newCustomerName == "" {
			return fmt.Errorf("Error : customer name cannot be empty")
		}
		if customerId == "" {
			return fmt.Errorf("Error : customer id cannot be empty")
		}

		updateCustomerError := dbHandler.Execute(`
      UPDATE customer SET customer_name = ? , customer_information = ?
      WHERE customer_id = ?
    `, newCustomerName, newCustomerInformation, customerId)
		if updateCustomerError != nil {
			return updateCustomerError
		}

		return nil

	}
}
