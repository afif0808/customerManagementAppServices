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
		dbHandler.Execute(`
      INSERT INTO customers(customer_name)VALUES(?);
    `, customerName)
		dbHandler.Execute(`
		INSERT INTO customers_information(customer_id,customer_information)VALUES(LAST_INSERT_ID(),?)
	`, customerInformation)
		return nil
	}
}
