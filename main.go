package main

import (
	"github.com/furqonzt99/todo-api/configs"
	projectContoller "github.com/furqonzt99/todo-api/delivery/controllers/project"
	todoContoller "github.com/furqonzt99/todo-api/delivery/controllers/todo"
	userContoller "github.com/furqonzt99/todo-api/delivery/controllers/user"
	"github.com/furqonzt99/todo-api/delivery/routes"
	projectRepo "github.com/furqonzt99/todo-api/repository/project"
	todoRepo "github.com/furqonzt99/todo-api/repository/todo"
	userRepo "github.com/furqonzt99/todo-api/repository/user"
	"github.com/furqonzt99/todo-api/utils"
	"github.com/labstack/echo/v4"
)

func main()  {
	config := configs.GetConfig()

	db := utils.InitDB(config)
	
	utils.InitialMigrate(db)

	userRepo := userRepo.NewUserRepo(db)
	userContoller := userContoller.NewUserController(userRepo)
  
	todoRepo := todoRepo.NewTodoRepo(db)
	todoContoller := todoContoller.NewTodoController(todoRepo)
	
	projectRepo := projectRepo.NewProjectRepo(db)
	projectContoller := projectContoller.NewProjectController(projectRepo)

	e := echo.New()

	routes.RegisterTodoPath(e, *todoContoller)
  	routes.RegisterUserPath(e, *userContoller)
  	routes.RegisterProjectPath(e, *projectContoller)

	e.Logger.Fatal(e.Start(":" + config.Port))
}
