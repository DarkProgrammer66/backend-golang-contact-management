package model

type User struct {
	Username string `gorm:"primaryKey;type:varchar(100)"`
	Password string `gorm:"type:varchar(255);not null"`
	Name     string `gorm:"type:varchar(255);not null"`
	Token    string `gorm:"type:text;default:null"`
}
