package models

import "gorm.io/gorm"

type LinkModel struct {
	gorm.Model

	Url  string `json:"url"`
	Name string `json:"name"`
}
