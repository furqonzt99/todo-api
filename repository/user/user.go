package user

import (
	"github.com/furqonzt99/todo-api/models"
	"gorm.io/gorm"
)

type Repository interface {
	Register(u models.User) (models.User, error)
	Login(name, pwd string) (models.User, error)

	GetUser(id int) (models.User, error)
	Delete(id int) error
	Update(user models.User) (models.User, error)
	GetLoginData(email string) (models.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) Register(user models.User) (models.User, error) {
	err := ur.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, err
}

func (ur *UserRepository) Login(name, pwd string) (models.User, error) {
	login := models.User{}
	if err := ur.db.Where("email = ? AND password = ?", name, pwd).Find(&login).Error; err != nil {
		return login, err
	}
	return login, nil
}

func (ur *UserRepository) GetUser(id int) (models.User, error) {
	var user models.User
	err := ur.db.First(&user, id).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (ur *UserRepository) Delete(id int) error {
	var user models.User
	err := ur.db.Unscoped().Delete(&user, id).Error
	if err != nil {
		return err
	}
	return err
}

func (ur *UserRepository) Update(user models.User) (models.User, error) {
	err := ur.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, err
}

func (ur *UserRepository) GetLoginData(email string) (models.User, error) {
	var user models.User
	err := ur.db.First(&user, "email = ?", email).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
