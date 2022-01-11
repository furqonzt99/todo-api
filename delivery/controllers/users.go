package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/furqonzt99/todo-api/delivery/middlewares"
	"github.com/furqonzt99/todo-api/models"
	"github.com/furqonzt99/todo-api/repository"
	"github.com/labstack/echo/v4"
)

type User struct {
	repository repository.Repository
}

func NewUsers(repository repository.Repository) *User {
	return &User{repository: repository}
}

func (u User) Register(c echo.Context) error {
	var user models.User
	c.Bind(&user)

	hash, _ := middlewares.Hashpwd(user.Password)

	user.Password = hash

	res, err := u.repository.Register(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(200, map[string]interface{}{
		"messages": "register success",
		"user":     res,
	})
}

func (u User) Login(c echo.Context) error {
	var login models.User
	c.Bind(&login)

	user, err := u.repository.GetLoginData(login.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("email %v not found!", login.Email))
	}

	hash, err := middlewares.Checkpwd(user.Password, login.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
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

func (u User) GetAll(c echo.Context) error {

	users, err := u.repository.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(200, map[string]interface{}{
		"messages": "success find all users",
		"users":    users,
	})

}

func (u User) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	idUser, err := u.repository.GetUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("id number %v not found!", id))
	}

	return c.JSON(200, map[string]interface{}{
		"messages": fmt.Sprintf("id %v has been found!", id),
		"user":     idUser,
	})
}

func (u User) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err2 := u.repository.Delete(id)
	if err2 != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(200, map[string]interface{}{
		"messages": fmt.Sprintf("id %v has been deleted!", id),
	})
}

func (u User) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user, err := u.repository.GetUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("id number %v not found!", id))
	}

	var tmpUser models.User
	c.Bind((&tmpUser))
	user.Name = tmpUser.Name
	user.Email = tmpUser.Email
	user.Password = tmpUser.Password

	userRes, err := u.repository.Update(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": fmt.Sprintf("id %v was updated!", id),
		"user":     userRes,
	})
}
