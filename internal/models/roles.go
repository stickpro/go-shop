package models

import "time"

type Roles struct {
	Id        uint
	Name      string `json:"name" gorm:"unique; not null"`
	CreatedAt time.Time
}

type RoleUser struct {
}
