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
      UPDATE customers SET customer_name = ?
      WHERE customer_id = ?
    `, newCustomerName, customerId)
		if updateCustomerError != nil {
			return updateCustomerError
		}

		updateCustomerInforamtionError := dbHandler.Execute(`
      UPDATE customers_information SET customer_information = ?
      WHERE customer_id = ?
  `, newCustomerInformation, customerId)

		if updateCustomerInforamtionError != nil {
			return updateCustomerInforamtionError
		}
		return nil

	}
}
