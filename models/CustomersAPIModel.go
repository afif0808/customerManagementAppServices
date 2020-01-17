package models


type BulkCustomersAPIModel struct {
	NextLink     string                 `json:"next"`
	PreviousLink string                 `json:"previous"`
	Count        string                 `json:"count"`
	Result       CustomerModel `json:"result"`
}

type SingleCustomerAPIModel struct {
	Result models.CustomerModel `json:"result"`
}
