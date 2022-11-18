package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Name string    
	MobileNumber string 
}
type Description struct {
	gorm.Model
	userId string
	title string
	body string
}

		