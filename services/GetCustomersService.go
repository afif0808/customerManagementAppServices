package services

import (
	"customerManagementAppServices/interfaces"
	"customerManagementAppServices/models"
)

func GetCustomerService(dbHandler interfaces.IDBHandler) models.GetCustomersServiceModel {

	return models.GetCustomersServiceModel(func(limit int, afterId int) ([]models.CustomerModel, error) {

		return nil, nil
	})
}
