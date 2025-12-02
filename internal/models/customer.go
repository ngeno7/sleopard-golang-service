package models

type Customer struct {
	Id int64 `json:"id"`
	Phone string `json:"phone"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Location string `json:"location"`
	PreferredProduct string `json:"preferred_product"`
}
