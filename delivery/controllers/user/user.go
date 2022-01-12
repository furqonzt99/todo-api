package user

import (
	"fmt"
	"net/http"

	"github.com/furqonzt99/todo-api/delivery/middlewares"
	"github.com/furqonzt99/todo-api/models"
	"github.com/furqonzt99/todo-api/repository/user"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	Repo user.Repository
}

func NewUserController(user user.Repository) *UserController {
	return &UserController{Repo: user}
}

func (uc UserController) Register(c echo.Context) error {
	var user models.User
	c.Bind(&user)

	hash, _ := middlewares.Hashpwd(user.Password)

	user.Password = hash

	res, err := uc.Repo.Register(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(200, map[string]interface{}{
		"messages": "register success",
		"user":     res,
	})
}

func (uc UserController) Login(c echo.Context) error {
	var login models.User
	c.Bind(&login)

	user, err := uc.Repo.GetLoginData(login.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("email %v not found!", login.Email))
	}

	hash, err := middlewares.Checkpwd(user.Password, login.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "password salah!")
	}

	var token string
	if hash {
		token, err = middlewares.CreateToken(int(user.ID))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
	}

	return c.JSON(200, map[string]interface{}{
		"messeges": "login succes",
		"token":    token,
	})
}

func (uc UserController) GetUser(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	idUser, err := uc.Repo.GetUser(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("id number %v not found!", userId))
	}

	return c.JSON(200, map[string]interface{}{
		"messages": fmt.Sprintf("id %v has been found!", userId),
		"user":     idUser,
	})
}

func (uc UserController) Delete(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	err := uc.Repo.Delete(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(200, map[string]interface{}{
		"messages": fmt.Sprintf("id %v has been deleted!", userId),
	})
}

func (uc UserController) Update(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	user, err := uc.Repo.GetUser(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("id number %v not found!", userId))
	}

	var tmpUser models.User
	c.Bind((&tmpUser))
	user.Name = tmpUser.Name
	user.Email = tmpUser.Email
	user.Password = tmpUser.Password

	userRes, err := uc.Repo.Update(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": fmt.Sprintf("id %v was updated!", userId),
		"user":     userRes,
	})
}
