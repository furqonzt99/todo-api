package routers

import (
	"github.com/furqonzt99/todo-api/dellivery/controllers"
	"github.com/labstack/echo/v4"
)

func Users(e *echo.Echo, cc *controllers.Users) {

	e.POST("/register", cc.Register)
	e.POST("/login", cc.Login)

}
