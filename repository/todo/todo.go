package todo

import (
	"github.com/furqonzt99/todo-api/models"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepo(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

type Todo interface {
	GetAll(userId int) ([]models.Todo, error)
	Get(userId int, todoId int) (models.Todo, error)
	Insert(models.Todo) (models.Todo, error)
	Edit(userId int, todoId int, todo models.Todo) (models.Todo, error)
	Delete(userId int, todoId int) (models.Todo, error)
}

func (tr *TodoRepository) GetAll(userId int) ([]models.Todo, error) {
	var todos []models.Todo

	if err := tr.db.Where("user_id = ?", userId).Find(&todos).Error; err != nil {
		return nil, err
	}

	return todos, nil
}

func (tr *TodoRepository) Get(userId int, todoId int) (models.Todo, error) {
	var todo models.Todo

	if err := tr.db.Where("user_id = ?", userId).First(&todo, todoId, "user_id = ?", userId).Error; err != nil {
		return todo, err
	}

	return todo, nil
}

func (tr *TodoRepository) Insert(todo models.Todo) (models.Todo, error) {

	if err := tr.db.Create(&todo).Error; err != nil {
		return todo, err
	}

	return todo, nil
}

func (tr *TodoRepository) Edit(userId int, todoId int, todo models.Todo) (models.Todo, error) {
	var t models.Todo
	tr.db.Where("user_id = ?", userId).First(&t, todoId)

	if err := tr.db.Model(&t).Updates(todo).Error; err != nil {
		return t, err
	}

	return t, nil
}

func (tr *TodoRepository) Delete(userId int, todoId int) (models.Todo, error) {
	var todo models.Todo

	if err := tr.db.Where("user_id = ?", userId).Delete(&todo, todoId).Error; err != nil {
		return todo, err
	}

	return todo, nil
}