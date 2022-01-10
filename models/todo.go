package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Name string `json:"name" form:"name"`
	TodoStart time.Time `json:"todo_start" form:"todo_start"`
	TodoEnd time.Time `json:"todo_end" form:"todo_end"`
	Status string `json:"status" form:"status"`
	UserID uint
	ProjectID uint `json:"project_id" form:"project_id"`
}