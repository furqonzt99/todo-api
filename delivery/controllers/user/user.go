package user

import (
	"net/http"

	"github.com/furqonzt99/todo-api/delivery/common"
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

	if err := c.Validate(user); err != nil {
      return c.JSON(http.StatusBadRequest, common.ErrorResponse(http.StatusBadRequest, err.Error()))
    }

	hash, _ := middlewares.Hashpwd(user.Password)

	user.Password = hash

	res, err := uc.Repo.Register(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ErrorResponse(http.StatusBadRequest, "email already exist"))
	}

	return c.JSON(http.StatusOK, common.SuccessResponse(res))
}

func (uc UserController) Login(c echo.Context) error {
	var login models.User
	c.Bind(&login)

	if err := c.Validate(login); err != nil {
      return c.JSON(http.StatusBadRequest, common.ErrorResponse(http.StatusBadRequest, err.Error()))
    }

	user, err := uc.Repo.GetLoginData(login.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, common.ErrorResponse(http.StatusNotFound, "not registered"))
	}

	hash, err := middlewares.Checkpwd(user.Password, login.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ErrorResponse(http.StatusBadRequest, "wrong password"))
	}

	var token string

	if hash {
		token, _ = middlewares.CreateToken(int(user.ID))
	}

	return c.JSON(http.StatusOK, common.SuccessResponse(token))
}

func (uc UserController) GetUser(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	user, err := uc.Repo.GetUser(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, common.ErrorResponse(http.StatusNotFound, "not found"))
	}

	return c.JSON(http.StatusOK, common.SuccessResponse(user))
}

func (uc UserController) Delete(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	err := uc.Repo.Delete(userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, common.ErrorResponse(http.StatusNotFound, err.Error()))
	}

	return c.JSON(http.StatusOK, common.SuccessResponse(err))
}

func (uc UserController) Update(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	user, err := uc.Repo.GetUser(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, common.ErrorResponse(http.StatusNotFound, "not found"))
	}

	var tmpUser models.User
	c.Bind((&tmpUser))
	user.Name = tmpUser.Name
	user.Email = tmpUser.Email

	hash, _ := middlewares.Hashpwd(user.Password)

	user.Password = hash

	if err := c.Validate(user); err != nil {
      return c.JSON(http.StatusBadRequest, common.ErrorResponse(http.StatusBadRequest, err.Error()))
    }

	userRes, err := uc.Repo.Update(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusOK, common.SuccessResponse(userRes))
}
