package models

type BulkCustomersAPIModel struct {
	NextPageLink       string          `json:"next"`
	PreviousPageLink   string          `json:"previous"`
	CustomerLimitCount int             `json:"count"`
	Result             []CustomerModel `json:"result"`
}

type SingleCustomerAPIModel struct {
	Result CustomerModel `json:"result"`
}
