package models

// GetCustomersServiceModel returns to at most 'limit' customers starting from customer number 'offset'
// the customer list is sorted by id
type GetCustomersModel func(limit, offset int) ([]CustomerModel, error)

// SearchCustomersServiceModel returns to at most 'limit' customers starting from starting from customer number 'offset'
// starting from the 'offset' customer number
// it searches by customer name
// the customer list is sorted by id
type SearchCustomersModel func(limit, offset int, keyword string) ([]CustomerModel, error)

type GetLastCustomerIdModel func() (int, error)
type GetLastCustomerIdInSearchModel func(keyword string) (int, error)
type GetSingleCustomerById func(id int) (*CustomerModel, error)
