package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/furqonzt99/todo-api/configs"
	"github.com/furqonzt99/todo-api/constants"
	"github.com/furqonzt99/todo-api/delivery/common"
	userContoller "github.com/furqonzt99/todo-api/delivery/controllers/user"
	"github.com/furqonzt99/todo-api/delivery/middlewares"
	userRepo "github.com/furqonzt99/todo-api/repository/user"
	"github.com/furqonzt99/todo-api/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

var token string

func TestRegisterUser(t *testing.T)  {
	config := configs.GetConfig()

	db := utils.InitDB(config)
	utils.InitialMigrate(db)

	userRepo := userRepo.NewUserRepo(db)
	userContoller := userContoller.NewUserController(userRepo)
	
	e := echo.New()
	
	t.Run("Register Success 1", func(t *testing.T) {
		e.POST("/register", userContoller.Register)

		registerBody, _ := json.Marshal(map[string]interface{}{
			"name": "Furqon",
			"email": "furqonzt99@gmail.com",
			"password": "1234qwer",
		})

		req := httptest.NewRequest(echo.POST, "/register", bytes.NewBuffer(registerBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		var response common.ResponseSuccess

		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "success", response.Message)
		assert.Equal(t, "Furqon", response.Data.(map[string]interface{})["name"])
		assert.Equal(t, "furqonzt99@gmail.com", response.Data.(map[string]interface{})["email"])
		assert.NotNil(t, response.Data.(map[string]interface{})["password"])

	})
	
	t.Run("Register Success 2", func(t *testing.T) {
		e.POST("/register", userContoller.Register)

		registerBody, _ := json.Marshal(map[string]interface{}{
			"name": "Furqon Nih",
			"email": "furqonzt98@gmail.com",
			"password": "1234qwer",
		})

		req := httptest.NewRequest(echo.POST, "/register", bytes.NewBuffer(registerBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		var response common.ResponseSuccess

		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "success", response.Message)
		assert.Equal(t, "Furqon Nih", response.Data.(map[string]interface{})["name"])
		assert.Equal(t, "furqonzt98@gmail.com", response.Data.(map[string]interface{})["email"])
		assert.NotNil(t, response.Data.(map[string]interface{})["password"])

	})
	
	t.Run("Register Failed (Email Already Exist)", func(t *testing.T) {
		e.POST("/register", userContoller.Register)

		registerBody, _ := json.Marshal(map[string]interface{}{
			"name": "Furqon",
			"email": "furqonzt99@gmail.com",
			"password": "1234qwer",
		})

		req := httptest.NewRequest(echo.POST, "/register", bytes.NewBuffer(registerBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		var response common.ResponseSuccess

		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.NotNil(t, response.Message)
	})
}

func TestLoginUser(t *testing.T)  {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	userRepo := userRepo.NewUserRepo(db)
	userContoller := userContoller.NewUserController(userRepo)
	
	e := echo.New()
	
	t.Run("Login Success", func(t *testing.T) {
		e.POST("/login", userContoller.Login)

		loginBody, _ := json.Marshal(map[string]interface{}{
			"email": "furqonzt99@gmail.com",
			"password": "1234qwer",
		})

		req := httptest.NewRequest(echo.POST, "/login", bytes.NewBuffer(loginBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		var response common.ResponseSuccess

		json.Unmarshal(rec.Body.Bytes(), &response)

		token = response.Data.(string)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "success", response.Message)
		assert.NotNil(t, response.Data)

	})
	
	t.Run("Login Failed (User Not Found)", func(t *testing.T) {
		e.POST("/login", userContoller.Login)

		loginBody, _ := json.Marshal(map[string]interface{}{
			"email": "furqonzt98332@gmail.com",
			"password": "1234qwesdrdad",
		})

		req := httptest.NewRequest(echo.POST, "/login", bytes.NewBuffer(loginBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		var response common.ResponseSuccess

		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.NotNil(t, response.Message)
	})
	
	t.Run("Login Failed (Wrong Password)", func(t *testing.T) {
		e.POST("/login", userContoller.Login)

		loginBody, _ := json.Marshal(map[string]interface{}{
			"email": "furqonzt99@gmail.com",
			"password": "1234qwerui",
		})

		req := httptest.NewRequest(echo.POST, "/login", bytes.NewBuffer(loginBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		var response common.ResponseSuccess

		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.NotNil(t, response.Message)
	})
}

func TestGetUser(t *testing.T)  {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	userRepo := userRepo.NewUserRepo(db)
	userContoller := userContoller.NewUserController(userRepo)
	
	e := echo.New()
	
	t.Run("Get User Success", func(t *testing.T) {
		e.GET("/users/profile", userContoller.GetUser, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		req := httptest.NewRequest(echo.GET, "/users/profile", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		var response common.ResponseSuccess

		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "success", response.Message)
		assert.Equal(t, "Furqon", response.Data.(map[string]interface{})["name"])
		assert.Equal(t, "furqonzt99@gmail.com", response.Data.(map[string]interface{})["email"])
		assert.NotNil(t, response.Data.(map[string]interface{})["password"])
	})
	
	t.Run("Get User Failed Not Found", func(t *testing.T) {
		e.GET("/users/profile", userContoller.GetUser, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		wrongToken, _ := middlewares.CreateToken(4)

		req := httptest.NewRequest(echo.GET, "/users/profile", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", wrongToken))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.NotNil(t, rec.Body)
	})

	t.Run("Get User Failed Unauthorize", func(t *testing.T) {
		e.GET("/users/profile", userContoller.GetUser, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		req := httptest.NewRequest(echo.GET, "/users/profile", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token + "wrongtoken"))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.NotNil(t, rec.Body)
	})
}

func TestUpdateUser(t *testing.T)  {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	userRepo := userRepo.NewUserRepo(db)
	userContoller := userContoller.NewUserController(userRepo)
	
	e := echo.New()
	
	t.Run("Update User Success", func(t *testing.T) {
		e.PUT("/users/update", userContoller.Update, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		dataBody, _ := json.Marshal(map[string]interface{}{
			"name": "Arif Furqon",
			"email": "furqonzt99@gmail.com",
			"password": "1234qwer",
		})

		req := httptest.NewRequest(echo.PUT, "/users/update", bytes.NewBuffer(dataBody))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		var response common.ResponseSuccess

		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "success", response.Message)
		assert.Equal(t, "Arif Furqon", response.Data.(map[string]interface{})["name"])
		assert.Equal(t, "furqonzt99@gmail.com", response.Data.(map[string]interface{})["email"])
		assert.NotNil(t, response.Data.(map[string]interface{})["password"])
	})
	
	t.Run("Update User Failed Not Found", func(t *testing.T) {
		e.PUT("/users/update", userContoller.Update, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		dataBody, _ := json.Marshal(map[string]interface{}{
			"name": "Arif Furqon",
			"email": "furqonzt99@gmail.com",
			"password": "1234qwer",
		})

		wrongToken, _ := middlewares.CreateToken(5)

		req := httptest.NewRequest(echo.PUT, "/users/update", bytes.NewBuffer(dataBody))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", wrongToken))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.NotNil(t, rec.Body)
	})
	
	t.Run("Update User Failed Bad Request", func(t *testing.T) {
		e.PUT("/users/update", userContoller.Update, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		dataBody, _ := json.Marshal(map[string]interface{}{
			"name": "Arif Furqon",
			"email": "furqonzt98@gmail.com",
			"password": "1234qwer",
		})

		req := httptest.NewRequest(echo.PUT, "/users/update", bytes.NewBuffer(dataBody))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.NotNil(t, rec.Body)
	})
	
	t.Run("Update User Failed Unauthorize", func(t *testing.T) {
		e.PUT("/users/update", userContoller.Update, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		dataBody, _ := json.Marshal(map[string]interface{}{
			"name": "Arif Furqon",
			"email": "furqonzt99@gmail.com",
			"password": "1234qwer",
		})

		req := httptest.NewRequest(echo.PUT, "/users/update", bytes.NewBuffer(dataBody))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token + "wrongtoken"))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.NotNil(t, rec.Body)
	})
}

func TestDeleteUser(t *testing.T)  {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	userRepo := userRepo.NewUserRepo(db)
	userContoller := userContoller.NewUserController(userRepo)
	
	e := echo.New()
	
	t.Run("Delete User Success", func(t *testing.T) {
		e.DELETE("/users/delete", userContoller.GetUser, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		req := httptest.NewRequest(echo.DELETE, "/users/delete", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		var response common.ResponseSuccess

		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "success", response.Message)
	})
	
	t.Run("Delete User Failed Not Found", func(t *testing.T) {
		e.DELETE("/users/delete", userContoller.GetUser, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		wrongToken, _ := middlewares.CreateToken(17)

		req := httptest.NewRequest(echo.DELETE, "/users/delete", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", wrongToken))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.NotNil(t, rec.Body)
	})

	t.Run("Delete User Failed Unauthorize", func(t *testing.T) {
		e.DELETE("/users/delete", userContoller.GetUser, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		req := httptest.NewRequest(echo.DELETE, "/users/delete", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token + "wrongtoken"))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.NotNil(t, rec.Body)
	})
}