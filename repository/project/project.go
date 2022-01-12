package project

import (
	"github.com/furqonzt99/todo-api/models"
	"gorm.io/gorm"
)

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepo(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

type Project interface {
	GetAll(userId int) ([]models.Project, error)
	Get(userId int, projectId int) (models.Project, error)
	Insert(models.Project) (models.Project, error)
	Edit(userId int, projectId int, project models.Project) (models.Project, error)
	Delete(userId int, projectId int) (models.Project, error)
}

func (tr *ProjectRepository) GetAll(userId int) ([]models.Project, error) {
	var projects []models.Project

	if err := tr.db.Preload("Todo").Where("user_id = ?", userId).Find(&projects).Error; err != nil {
		return nil, err
	}

	return projects, nil
}

func (tr *ProjectRepository) Get(userId int, projectId int) (models.Project, error) {
	var project models.Project

	if err := tr.db.Preload("Todo").Where("user_id = ?", userId).First(&project, projectId).Error; err != nil {
		return project, err
	}

	return project, nil
}

func (tr *ProjectRepository) Insert(project models.Project) (models.Project, error) {

	if err := tr.db.Create(&project).Error; err != nil {
		return project, err
	}

	return project, nil
}

func (tr *ProjectRepository) Edit(userId int, projectId int, project models.Project) (models.Project, error) {
	var t models.Project
	tr.db.Where("user_id = ?", userId).First(&t, projectId)

	if err := tr.db.Model(&t).Updates(project).Error; err != nil {
		return t, err
	}

	return t, nil
}

func (tr *ProjectRepository) Delete(userId int, projectId int) (models.Project, error) {
	var project models.Project

	if err := tr.db.Where("user_id = ?", userId).Delete(&project, projectId).Error; err != nil {
		return project, err
	}

	return project, nil
}