package models

type CustomerModel struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Information string `json:"information"`
	DateAdded   string `json:"date_added"`
}
