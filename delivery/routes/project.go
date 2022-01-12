package routes

import (
	"github.com/furqonzt99/todo-api/constants"
	controller "github.com/furqonzt99/todo-api/delivery/controllers/project"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterProjectPath(e *echo.Echo, pc controller.ProjectContoller)  {

	auth := e.Group("")
	auth.Use(middleware.JWT([]byte(constants.JWT_SECRET_KEY)))
	
	auth.GET("/projects", pc.GetAll)
	auth.GET("/projects/:id", pc.Get)
	auth.POST("/projects", pc.Insert)
	auth.PUT("/projects/:id", pc.Edit)
	auth.DELETE("/projects/:id", pc.Delete)
}