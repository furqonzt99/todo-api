package todo

import (
	"net/http"
	"strconv"
	"time"

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

func (tc TodoContoller) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		todos, err := tc.Repo.GetAll()

		if err != nil {
			return c.JSON(http.StatusNotFound, err.Error())
		}

		return c.JSON(http.StatusOK, todos)
	}
}

func (tc TodoContoller) Get() echo.HandlerFunc{
	return func(c echo.Context) error {
		todoId, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		todo, err := tc.Repo.Get(todoId)

		if err != nil {
			return c.JSON(http.StatusNotFound, err.Error())
		}

		return c.JSON(http.StatusOK, todo)
	}
}

func (tc TodoContoller) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		todoData := models.Todo{}
		c.Bind(&todoData)

		loc, _ := time.LoadLocation("Asia/Jakarta")
		
		todoStart := todoData.TodoStart
		localTodoStart := time.Date(todoStart.Year(), todoStart.Month(), todoStart.Day(), todoStart.Hour(), todoStart.Minute(), todoStart.Second(), todoStart.Nanosecond(), loc)
		
		todoData.TodoStart = localTodoStart
		
		todoEnd := todoData.TodoEnd
		localTodoEnd := time.Date(todoEnd.Year(), todoEnd.Month(), todoEnd.Day(), todoEnd.Hour(), todoEnd.Minute(), todoEnd.Second(), todoEnd.Nanosecond(), loc)
		
		todoData.TodoEnd = localTodoEnd

		todo, err := tc.Repo.Insert(todoData)

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, todo)
	}
}

func (tc TodoContoller) Edit() echo.HandlerFunc {
	return func(c echo.Context) error {
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

		todo, err := tc.Repo.Edit(todoId, todoData)

		if err != nil {
			return c.JSON(http.StatusNotFound, err.Error())
		}

		return c.JSON(http.StatusOK, todo)
	}
}

func (tc TodoContoller) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		todoId, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		todo, err := tc.Repo.Delete(todoId)

		if err != nil {
			return c.JSON(http.StatusNotFound, err.Error())
		}

		return c.JSON(http.StatusOK, todo)
	}
}

