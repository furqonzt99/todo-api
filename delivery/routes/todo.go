package routes

import (
	"github.com/furqonzt99/todo-api/constants"
	controller "github.com/furqonzt99/todo-api/delivery/controllers/todo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterTodoPath(e *echo.Echo, tc controller.TodoContoller)  {

	auth := e.Group("")
	auth.Use(middleware.JWT([]byte(constants.JWT_SECRET_KEY)))
	
	auth.GET("/todos", tc.GetAll)
	auth.GET("/todos/:id", tc.Get)
	auth.POST("/todos", tc.Insert)
	auth.PUT("/todos/:id", tc.Edit)
	auth.DELETE("/todos/:id", tc.Delete)

	auth.POST("/todos/:id/complete", tc.SetComplete)
	auth.POST("/todos/:id/reopen", tc.Reopen)
}