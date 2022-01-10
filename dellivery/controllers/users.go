package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/furqonzt99/todo-api/models"
	"github.com/furqonzt99/todo-api/repository"
	"github.com/labstack/echo/v4"
)

type Users struct {
	repository repository.Repository
}

func NewUsers(repository repository.Repository) *Users {
	return &Users{repository: repository}
}

func (u Users) Register(c echo.Context) error {
	var user models.User
	c.Bind(&user)

	res, err := u.repository.Register(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(200, map[string]interface{}{
		"messages": "register success",
		"user":     res,
	})
}

func (u Users) Login(c echo.Context) error {
	var login models.User
	c.Bind(&login)

	res, err := u.repository.Login(login.Email, login.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "id or your password salah")
	}

	return c.JSON(200, map[string]interface{}{
		"messeges": "login succes",
		"token":    res.Token,
	})
}

func (u Users) GetAll(c echo.Context) error {

	users, err := u.repository.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(200, map[string]interface{}{
		"messages": "success find all users",
		"users":    users,
	})

}

func (u Users) GetUser(c echo.Context) error {
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

func (u Users) Delete(c echo.Context) error {
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

func (u Users) Update(c echo.Context) error {
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
