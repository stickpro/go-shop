package models

import "gorm.io/gorm"

// User ... User Database Model
type User struct {
	gorm.Model
	Email     string `json:"email" gorm:"unique;not null"`
	Password  string `json:"password" gorm:"not null"`
	FirstName string `json:"first_name" gorm:"varchar(255);not null"`
	LastName  string `json:"last_name" gorm:"varchar(255);not null"`
	Status    bool   `json:"status" gorm:"not null;default:false"`
}

//MoneyTransfer --- MoneyTransfer Struct
type MoneyTransfer struct {
	Receiver uint    `json:"receiver"`
	Giver    uint    `json:"giver"`
	Amount   float64 `json:"amount"`
}

type UserRole struct {
	UserId int `json:"user_id"`
	RoleId int `json:"role_id"`
}
