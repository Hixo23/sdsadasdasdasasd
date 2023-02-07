package models

import "gorm.io/gorm"

type LinkModel struct {
	gorm.Model

	Url  string `json:"url" binding:"required" gorm:"unique_index"`
	Name string `json:"name" binding:"required" gorm:"unique_index"`
}
