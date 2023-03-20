package models

import "gorm.io/gorm"
type User struct {
	gorm.Model

	FirstName	string	`gorm:"type:varchar(60);not null" json:"first_name"`
	LastName	string	`gorm:"type:varchar(60);not null" json:"last_name"`
	Email		string	`gorm:"type:varchar(255);not null;unique" json:"email"`
	Tasks		[]Task	`json:"tasks"`
}