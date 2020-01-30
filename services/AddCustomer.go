package services

import (
	"customerManagementAppServices/interfaces"
	"customerManagementAppServices/models"
	"fmt"
)

func AddCustomer(dbHandler interfaces.IDBHandler) models.AddCustomerModel {
	return func(customerName, customerInformation string) error {
		if customerName == "" {
			return fmt.Errorf("Error : customer name cannot be empty")
		}
		addCustomerError := dbHandler.Execute(`
      INSERT INTO customers(customer_name)VALUES(?);
    `, customerName)
		if addCustomerError != nil {
			return addCustomerError
		}

		addCustomerInformationError := dbHandler.Execute(`
		INSERT INTO customers_information(customer_id,customer_information)VALUES(LAST_INSERT_ID(),?)
	`, customerInformation)
		if addCustomerInformationError != nil {
			return addCustomerInformationError
		}

		return nil
	}
}
