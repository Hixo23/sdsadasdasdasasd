package models

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	Url  string `json:"url"`
	Name string `json:"name"`
}

type BindLink struct {
	Url  string `json:"url" binding:"required"`
	Name string `json:"name" binding:"required"`
}
