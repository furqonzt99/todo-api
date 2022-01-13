package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Name string `json:"name" form:"name" validate:"required"`
	TodoStart time.Time `json:"todo_start" form:"todo_start" validate:"required"`
	TodoEnd time.Time `json:"todo_end" form:"todo_end" validate:"required"`
	Status string `json:"status" form:"status" gorm:"default:'uncomplete'"`
	UserID uint `json:"user_id" validate:"required"`
	ProjectID uint `json:"project_id" form:"project_id" gorm:"default:null"`
}