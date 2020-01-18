package models

// GetCustomersServiceModel returns to at most 'limit' customers starting from customer number 'offset'
// the customer list is sorted by id
type GetCustomersServiceModel func(limit, afterId int) ([]CustomerModel, error)

// SearchCustomersServiceModel returns to at most 'limit' customers starting from starting from customer number 'offset'
// starting from the 'offset' customer number
// it searches by customer name
// the customer list is sorted by id
type SearchCustomersServiceModel func(limit, afterId int, keyword string) ([]CustomerModel, error)
