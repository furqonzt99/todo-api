package main

import (
	"github.com/furqonzt99/todo-api/configs"
	"github.com/furqonzt99/todo-api/delivery/controllers"
	"github.com/furqonzt99/todo-api/delivery/routers"
	"github.com/furqonzt99/todo-api/repository"
	"github.com/furqonzt99/todo-api/utils"
	"github.com/labstack/echo/v4"
)

func main() {
	config := configs.GetConfig()
	db := utils.InitDB(config)
	repoUser := repository.NewUser(db)
	ctrlUser := controllers.NewUsers(repoUser)

	e := echo.New()
	routers.User(e, ctrlUser)
	e.Logger.Fatal(e.Start(":8080"))
}
