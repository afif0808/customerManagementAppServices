package services

import (
	"customerManagementAppServices/interfaces"
	"customerManagementAppServices/models"
	"fmt"
	"log"
)

func AddCustomer(dbHandler interfaces.IDBHandler) models.AddCustomerModel {
	return func(customerName, customerInformation string) error {
		if customerName == "" {
			return fmt.Errorf("Error : customer name cannot be empty")
		}
		addCustomerError := dbHandler.Execute(`
      INSERT INTO customer(customer_name,customer_information)VALUES(?,?);
    `, customerName, customerInformation)
		if addCustomerError != nil {
			log.Println(addCustomerError)
			return addCustomerError
		}

		return nil
	}
}
