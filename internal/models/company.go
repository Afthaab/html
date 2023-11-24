package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name     string `json:"name" gorm:"unique" `
	Location string `json:"location" `
	Field    string `json:"field" `
}
