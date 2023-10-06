package model

import (
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Title   string `json:"title" binding:"min=2,max=100" gorm:"type:varchar(100)"`
	Content string `json:"content" binding:"max=100" gorm:"varchar(200)"`
}
