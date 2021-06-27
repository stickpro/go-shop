package models

import "time"

// User ... User Database Model
type User struct {
	ID        uint
	Email     string `json:"email" gorm:"unique;not null"`
	Password  string `json:"password" gorm:"not null"`
	FirstName string `json:"first_name" gorm:"varchar(255);not null"`
	LastName  string `json:"last_name" gorm:"varchar(255);not null"`
	Status    bool   `json:"status" gorm:"not null;default:false"`
	CreatedAt time.Time
	DeletedAt time.Time
}

type UserRole struct {
	UserId int `json:"user_id"`
	RoleId int `json:"role_id"`
}
