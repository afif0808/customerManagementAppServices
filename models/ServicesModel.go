package models

// GetCustomersFunc returns to at most 'limit' customers starting from customer number 'offset'
// sorted by id
type GetCustomersFunc func(limit, offset int) []CustomerModel

// SearchCustomersFunc returns the search results to at most 'limit' customers
// starting from the 'offset' customer number
// sorted by id
type SearchCustomersFunc func(limit, offset int, keyword string)
