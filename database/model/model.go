package model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	*gorm.Model
	Id        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type PlainModel struct {
	*gorm.Model
	Id uint `gorm:"primarykey" json:"id"`
}
