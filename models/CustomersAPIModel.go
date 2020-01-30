package models

type BulkCustomersAPIModel struct {
	NextPageLink     string          `json:"next"`
	PreviousPageLink string          `json:"previous"`
	Result           []CustomerModel `json:"result,null"`
}

type SingleCustomerAPIModel struct {
	Result *CustomerModel `json:"result"`
}
