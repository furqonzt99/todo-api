package todo

import (
	"net/http"
	"strconv"
	"time"

	"github.com/furqonzt99/todo-api/delivery/middlewares"
	"github.com/furqonzt99/todo-api/models"
	repository "github.com/furqonzt99/todo-api/repository/todo"
	"github.com/labstack/echo/v4"
)

type TodoContoller struct {
	Repo repository.Todo
}

func NewTodoController(todo repository.Todo) *TodoContoller {
	return &TodoContoller{Repo: todo}
}

func (tc TodoContoller) GetAll(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	todos, err := tc.Repo.GetAll(userId)

	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, todos)
}

func (tc TodoContoller) Get(c echo.Context) error {
	todoId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	userId := middlewares.ExtractTokenUserId(c)

	todo, err := tc.Repo.Get(userId, todoId)

	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, todo)
}

func (tc TodoContoller) Insert(c echo.Context) error {
	todoData := models.Todo{}
	c.Bind(&todoData)

	loc, _ := time.LoadLocation("Asia/Jakarta")
	
	todoStart := todoData.TodoStart
	localTodoStart := time.Date(todoStart.Year(), todoStart.Month(), todoStart.Day(), todoStart.Hour(), todoStart.Minute(), todoStart.Second(), todoStart.Nanosecond(), loc)
	
	todoData.TodoStart = localTodoStart
	
	todoEnd := todoData.TodoEnd
	localTodoEnd := time.Date(todoEnd.Year(), todoEnd.Month(), todoEnd.Day(), todoEnd.Hour(), todoEnd.Minute(), todoEnd.Second(), todoEnd.Nanosecond(), loc)
	
	todoData.TodoEnd = localTodoEnd

	userId := middlewares.ExtractTokenUserId(c)

	todoData.UserID = uint(userId)

	todo, err := tc.Repo.Insert(todoData)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, todo)
}

func (tc TodoContoller) Edit(c echo.Context) error {
	todoId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	todoData := models.Todo{}
	c.Bind(&todoData)

	loc, _ := time.LoadLocation("Asia/Jakarta")
	
	todoStart := todoData.TodoStart
	localTodoStart := time.Date(todoStart.Year(), todoStart.Month(), todoStart.Day(), todoStart.Hour(), todoStart.Minute(), todoStart.Second(), todoStart.Nanosecond(), loc)
	
	todoData.TodoStart = localTodoStart
	
	todoEnd := todoData.TodoEnd
	localTodoEnd := time.Date(todoEnd.Year(), todoEnd.Month(), todoEnd.Day(), todoEnd.Hour(), todoEnd.Minute(), todoEnd.Second(), todoEnd.Nanosecond(), loc)
	
	todoData.TodoEnd = localTodoEnd

	userId := middlewares.ExtractTokenUserId(c)

	todoData.UserID = uint(userId)

	todo, err := tc.Repo.Edit(userId, todoId, todoData)

	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, todo)
}

func (tc TodoContoller) Delete(c echo.Context) error {
	todoId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	userId := middlewares.ExtractTokenUserId(c)

	todo, err := tc.Repo.Delete(userId, todoId)

	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, todo)
}

func (tc TodoContoller) SetComplete(c echo.Context) error {
	todoId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	todoData := models.Todo{}
	c.Bind(&todoData)
	
	userId := middlewares.ExtractTokenUserId(c)
	
	todoData.UserID = uint(userId)
	todoData.Status = "complete"

	todo, err := tc.Repo.Edit(userId, todoId, todoData)

	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, todo)
}

func (tc TodoContoller) Reopen(c echo.Context) error {
	todoId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	todoData := models.Todo{}
	c.Bind(&todoData)
	
	userId := middlewares.ExtractTokenUserId(c)
	
	todoData.UserID = uint(userId)
	todoData.Status = "uncomplete"

	todo, err := tc.Repo.Edit(userId, todoId, todoData)

	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, todo)
}

