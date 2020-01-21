package models

type BulkCustomersAPIModel struct {
	NextPageLink     string `json:"next,omitempty"`
	PreviousPageLink string `json:"previous,omitempty"`
	// CustomerLimitCount int             `json:"count"`
	Result []CustomerModel `json:"result,null"`
}

type SingleCustomerAPIModel struct {
	Result CustomerModel `json:"result"`
}
