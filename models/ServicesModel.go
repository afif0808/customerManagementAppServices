package models

// GetCustomersFunc returns to at most 'limit' customers starting from 'afterId' customer
// the customer list is sorted by id
type GetCustomersFunc func(limit, afterId int) ([]CustomerModel, error)

// GetCustomersFunc returns to at most 'limit' customers starting from 'afterId' customer
// starting from the 'offset' customer number
// it searches by customer name
// the customer list is sorted by id
type SearchCustomersFunc func(limit, afterId int, keyword string) ([]CustomerModel, error)
