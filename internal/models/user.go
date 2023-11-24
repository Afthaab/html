package models

import "gorm.io/gorm"

type NewUser struct {
	Username    string `json:"username" validate:"required"`
	Email       string `json:"email" validate:"required"`
	Password    string `json:"password" validate:"required"`
	DateOfBirth string `json:"dateOfBirth" validate:"required"`
}

type User struct {
	gorm.Model
	Username     string `json:"username" gorm:"unique"`
	Email        string `json:"email" gorm:"unique"`
	Dob          string `json:"dob"`
	PasswordHash string `json:"-"`
}
