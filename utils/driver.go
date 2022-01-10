package utils

import (
	"fmt"

	"github.com/furqonzt99/todo-api/configs"
	"github.com/furqonzt99/todo-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *configs.AppConfig) *gorm.DB {

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", 
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)

	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

func InitialMigrate(db *gorm.DB)  {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Todo{})
	db.AutoMigrate(&models.Project{})
}