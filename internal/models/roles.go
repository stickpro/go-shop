package models

type Role struct {
	Id   uint   `gorm:"primaryKey"`
	Name string `json:"name" gorm:"unique; not null"`
}

type RoleUser struct {
	RoleId uint
	Role   Role
	UserId uint
	User   User
}
