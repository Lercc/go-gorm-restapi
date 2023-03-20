package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	UserId     		uint
	Title 			string	`gorm:"type:varchar(100);not null; unique_index"`
	Description  	string	`gorm:"type:text"`
	Done     		bool	`gorm:"default:false"`
}