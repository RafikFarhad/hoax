package model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	*gorm.Model
	Id        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_id"`
	UpdatedAt time.Time      `json:"updated_id"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type PlainModel struct {
	*gorm.Model
	Id uint `gorm:"primarykey" json:"id"`
}
