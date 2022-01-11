package routers

import (
	"github.com/furqonzt99/todo-api/constants"
	"github.com/furqonzt99/todo-api/delivery/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func User(e *echo.Echo, c *controllers.User) {

	auth := e.Group("")
	auth.Use(middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

	e.POST("/register", c.Register)
	e.POST("/login", c.Login)

	auth.GET("/users", c.GetAll)
	auth.GET("/users/:id", c.GetUser)
	auth.DELETE("/users/:id", c.Delete)
	auth.PUT("/users/:id", c.Update)
}
