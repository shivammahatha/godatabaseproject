package initializers

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string    `gorm:"size:255;not null;unique" json:"name"`
	MobileNumber     time.Time `gorm:"size:255;not null;unique" json:"mobilenumber"`
	
}
