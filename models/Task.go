package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	UserId     		uint	`json:"user_id"`
	Title 			string	`gorm:"type:varchar(100);not null; unique_index" json:"title"`
	Description  	string	`gorm:"type:text" json:"description"`
	Done     		bool	`gorm:"default:false" json:"done"`
}