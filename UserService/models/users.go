package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"type:varchar(100);unique_index"`
	Password string
}

func (u User) String() string {
	return fmt.Sprintf("%s - %s", u.Name, u.Email)
}
