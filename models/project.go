package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	UserID uint `json:"user_id"`
	Todo []Todo `json:"todo"`
}