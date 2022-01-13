package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name string `json:"name" form:"name" validate:"required"`
	Description string `json:"description" form:"description"`
	UserID uint `json:"user_id" validate:"required"`
	Todo []Todo `json:"todo"`
}