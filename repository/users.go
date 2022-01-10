package repository

import (
	"github.com/furqonzt99/todo-api/models"
	"gorm.io/gorm"
)

type Repository interface {
	Register(u models.User) (models.User, error)
	Login(name, pwd string) (models.User, error)

	GetAll() ([]models.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Register(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, err
}

func (u *repository) Login(name, pwd string) (models.User, error) {
	login := models.User{}
	if err := u.db.Where("email = ? AND password = ?", name, pwd).Find(&login).Error; err != nil {
		return login, err
	}
	return login, nil
}
func (u *repository) GetAll() ([]models.User, error) {
	var users []models.User
	err := u.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
