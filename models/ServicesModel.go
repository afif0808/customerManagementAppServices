package models

// GetCustomersFunc returns to at most 'limit' customers starting from 'afterId' customer
// the customer list is sorted by id
type GetCustomersFunc func() []CustomerModel

// GetCustomersFunc returns to at most 'limit' customers starting from 'afterId' customer
// starting from the 'offset' customer number
// it searches by customer name
// the customer list is sorted by id
type SearchCustomersFunc func(keyword string) []CustomerModel
