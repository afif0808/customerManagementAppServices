package models

//GetCustomersFunc returns 'limit' count customers starting from row 'offset'
type GetCustomersFunc func(limit, offset int) []CustomerModel
