package model

import (
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Username string `gorm:"uniqueIndex;not null;size:30"`
	Email    string `gorm:"uniqueIndex;size:50"`
	Password string `gorm:"not null;size:50"`
}
