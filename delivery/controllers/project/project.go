package project

import (
	"net/http"
	"strconv"

	"github.com/furqonzt99/todo-api/delivery/middlewares"
	"github.com/furqonzt99/todo-api/models"
	repository "github.com/furqonzt99/todo-api/repository/project"
	"github.com/labstack/echo/v4"
)

type ProjectContoller struct {
	Repo repository.Project
}

func NewProjectController(project repository.Project) *ProjectContoller {
	return &ProjectContoller{Repo: project}
}

func (tc ProjectContoller) GetAll(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	projects, err := tc.Repo.GetAll(userId)

	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, projects)
}

func (tc ProjectContoller) Get(c echo.Context) error {
	projectId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	userId := middlewares.ExtractTokenUserId(c)

	project, err := tc.Repo.Get(userId, projectId)

	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, project)
}

func (tc ProjectContoller) Insert(c echo.Context) error {
	projectData := models.Project{}
	c.Bind(&projectData)

	userId := middlewares.ExtractTokenUserId(c)

	projectData.UserID = uint(userId)

	project, err := tc.Repo.Insert(projectData)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, project)
}

func (tc ProjectContoller) Edit(c echo.Context) error {
	projectId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	projectData := models.Project{}
	c.Bind(&projectData)

	userId := middlewares.ExtractTokenUserId(c)

	projectData.UserID = uint(userId)

	project, err := tc.Repo.Edit(userId, projectId, projectData)

	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, project)
}

func (tc ProjectContoller) Delete(c echo.Context) error {
	projectId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	userId := middlewares.ExtractTokenUserId(c)

	project, err := tc.Repo.Delete(userId, projectId)

	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, project)
}