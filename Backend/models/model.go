package models

import "gorm.io/gorm"

type LinkModel struct {
	gorm.Model

	Url  string `json:"url" binding:"required,gte=4,lte=30" gorm:"unique_index"`
	Name string `json:"name" binding:"required,gte=1,lte=15" gorm:"unique_index"`
}
