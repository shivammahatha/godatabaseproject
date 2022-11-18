package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Name string    
	MobileNumber string 
}
type Apipost struct {
	gorm.Model
	UserId int
	Title string
	Body string
}

		