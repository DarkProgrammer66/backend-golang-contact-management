package model

type Contact struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	FirstName string `json:"first_name" gorm:"size:100;not null"`
	LastName  string `json:"last_name" gorm:"size:100"`
	Email     string `json:"email" gorm:"size:200"`
	Phone     string `json:"phone" gorm:"size:20"`
	Username  string `json:"username" gorm:"size:100;not null"`
}
