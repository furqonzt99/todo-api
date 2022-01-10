package routers

import (
	"github.com/furqonzt99/todo-api/dellivery/controllers"
	"github.com/labstack/echo/v4"
)

func Users(e *echo.Echo, cc *controllers.Users) {

	e.POST("/register", cc.Register)
	e.POST("/login", cc.Login)

	e.GET("/getAll", cc.GetAll)
	e.GET("/getuser/:id", cc.GetUser)
	e.DELETE("/delete/:id", cc.Delete)
	e.PUT("/update/:id", cc.Update)
}
