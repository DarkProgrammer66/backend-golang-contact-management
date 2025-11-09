package model

type User struct {
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Token    string `json:"token"`
}
