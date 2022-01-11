package main

import (
	"github.com/furqonzt99/todo-api/configs"
	todoContoller "github.com/furqonzt99/todo-api/delivery/controllers/todo"
	"github.com/furqonzt99/todo-api/delivery/routes"
	todoRepo "github.com/furqonzt99/todo-api/repository/todo"
	"github.com/furqonzt99/todo-api/utils"
	"github.com/labstack/echo/v4"
)

func main()  {
	config := configs.GetConfig()

	db := utils.InitDB(config)
	
	utils.InitialMigrate(db)

	todoRepo := todoRepo.NewTodoRepo(db)
	todoContoller := todoContoller.NewTodoController(todoRepo)

	e := echo.New()

	routes.RegisterTodoPath(e, *todoContoller)

	e.Logger.Fatal(e.Start(":" + config.Port))
}