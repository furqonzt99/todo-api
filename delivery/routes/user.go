package routes

import (
	"github.com/furqonzt99/todo-api/constants"
	controller "github.com/furqonzt99/todo-api/delivery/controllers/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterUserPath(e *echo.Echo, uc controller.UserController) {

	auth := e.Group("")
	auth.Use(middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

	e.POST("/register", uc.Register)
	e.POST("/login", uc.Login)

	auth.GET("/users/profile", uc.GetUser)
	auth.DELETE("/users/delete", uc.Delete)
	auth.PUT("/users/update", uc.Update)
}
