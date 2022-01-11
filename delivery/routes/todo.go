package routes

import (
	controller "github.com/furqonzt99/todo-api/delivery/controllers/todo"
	"github.com/labstack/echo/v4"
)

func RegisterTodoPath(e *echo.Echo, tc controller.TodoContoller)  {
	e.GET("/todos", tc.GetAll())
	e.GET("/todos/:id", tc.Get())
	e.POST("/todos", tc.Insert())
	e.PUT("/todos/:id", tc.Edit())
	e.DELETE("/todos/:id", tc.Delete())
}