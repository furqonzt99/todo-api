package repository

import (
	"github.com/furqonzt99/todo-api/models"
	"gorm.io/gorm"
)

type Repository interface {
	Register(u models.User) (models.User, error)
	Login(name, pwd string) (models.User, error)

	GetAll() ([]models.User, error)
	GetUser(id int) (models.User, error)
	Delete(id int) error
	Update(user models.User) (models.User, error)
	GetLoginData(email string) (models.User, error)
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

func (u *repository) GetUser(id int) (models.User, error) {
	var user models.User
	err := u.db.First(&user, id).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *repository) Delete(id int) error {
	var user models.User
	err := u.db.Delete(&user, id).Error
	if err != nil {
		return err
	}
	return err
}

func (u *repository) Update(user models.User) (models.User, error) {
	err := u.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, err
}

func (u *repository) GetLoginData(email string) (models.User, error) {
	var user models.User
	err := u.db.First(&user, "email = ?", email).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
