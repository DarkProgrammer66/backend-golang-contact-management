package model

type Address struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	ContactID  uint   `json:"contact_id"`
	Street     string `json:"street"`
	City       string `json:"city"`
	Province   string `json:"province"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code"`
}
