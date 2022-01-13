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
	todoContoller "github.com/furqonzt99/todo-api/delivery/controllers/todo"
	"github.com/furqonzt99/todo-api/delivery/middlewares"
	"github.com/furqonzt99/todo-api/models"
	todoRepo "github.com/furqonzt99/todo-api/repository/todo"
	"github.com/furqonzt99/todo-api/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

func TestInsertTodo(t *testing.T)  {
	config := configs.GetConfig()

	db := utils.InitDB(config)
	db.Migrator().DropTable(&models.Todo{})
	utils.InitialMigrate(db)

	todoRepo := todoRepo.NewTodoRepo(db)
	todoContoller := todoContoller.NewTodoController(todoRepo)
	
	e := echo.New()
	
	t.Run("Insert Todo Success", func(t *testing.T) {
		e.POST("/todos", todoContoller.Insert, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		todoBody, _ := json.Marshal(map[string]interface{}{
			"name": "Hiling",
			"todo_start": "2018-09-22T12:42:31+07:00",
			"todo_end": "2018-09-22T15:42:31+07:00",
		})

		req := httptest.NewRequest(echo.POST, "/todos", bytes.NewBuffer(todoBody))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		var response common.ResponseSuccess

		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "success", response.Message)
		assert.Equal(t, "Hiling", response.Data.(map[string]interface{})["name"])
		assert.Equal(t, "2018-09-22T12:42:31+07:00", response.Data.(map[string]interface{})["todo_start"])
		assert.Equal(t, "2018-09-22T15:42:31+07:00", response.Data.(map[string]interface{})["todo_end"])
		assert.Equal(t, "uncomplete", response.Data.(map[string]interface{})["status"])
	})
	
	t.Run("Insert Todo Failed Unauthorize", func(t *testing.T) {
		e.POST("/todos", todoContoller.Insert, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		todoBody, _ := json.Marshal(map[string]interface{}{
			"name": "Hiling",
			"todo_start": "2018-09-22T12:42:31+07:00",
			"todo_end": "2018-09-22T15:42:31+07:00",
		})

		req := httptest.NewRequest(echo.POST, "/todos", bytes.NewBuffer(todoBody))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token + "wrongtoken"))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		var response common.ResponseError

		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.NotNil(t, response.Message)
	})
}

func TestGetAllTodo(t *testing.T)  {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	todoRepo := todoRepo.NewTodoRepo(db)
	todoContoller := todoContoller.NewTodoController(todoRepo)
	
	e := echo.New()
	
	t.Run("Get All Todo Success", func(t *testing.T) {
		e.GET("/todos", todoContoller.GetAll, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		req := httptest.NewRequest(echo.GET, "/todos", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		var response common.ResponseSuccess

		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.NotNil(t, response.Data)
	})
	
	t.Run("Get Todo Failed Unauthorize", func(t *testing.T) {
		e.GET("/todos", todoContoller.GetAll, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		wrongToken, _ := middlewares.CreateToken(4)

		req := httptest.NewRequest(echo.GET, "/todos", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", wrongToken + "wrongtoken"))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.NotNil(t, rec.Body)
	})
}

func TestGetTodo(t *testing.T)  {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	todoRepo := todoRepo.NewTodoRepo(db)
	todoContoller := todoContoller.NewTodoController(todoRepo)
	
	e := echo.New()
	
	t.Run("Get Todo Success", func(t *testing.T) {
		e.GET("/todos/:id", todoContoller.Get, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		req := httptest.NewRequest(echo.GET, "/todos/1", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		var response common.ResponseSuccess

		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "Hiling", response.Data.(map[string]interface{})["name"])
		assert.Equal(t, "2018-09-22T12:42:31+07:00", response.Data.(map[string]interface{})["todo_start"])
		assert.Equal(t, "2018-09-22T15:42:31+07:00", response.Data.(map[string]interface{})["todo_end"])
		assert.Equal(t, "uncomplete", response.Data.(map[string]interface{})["status"])
	})
	
	t.Run("Get Todo Failed Not Found", func(t *testing.T) {
		e.GET("/todos/:id", todoContoller.Get, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		wrongToken, _ := middlewares.CreateToken(4)

		req := httptest.NewRequest(echo.GET, "/todos/4", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", wrongToken))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.NotNil(t, rec.Body)
	})
	
	t.Run("Get Todo Failed Unauthorize", func(t *testing.T) {
		e.GET("/todos/:id", todoContoller.Get, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		wrongToken, _ := middlewares.CreateToken(4)

		req := httptest.NewRequest(echo.GET, "/todos/1", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", wrongToken + "wrongtoken"))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.NotNil(t, rec.Body)
	})
}

func TestUpdateTodo(t *testing.T)  {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	todoRepo := todoRepo.NewTodoRepo(db)
	todoContoller := todoContoller.NewTodoController(todoRepo)
	
	e := echo.New()
	
	t.Run("Update Todo Success", func(t *testing.T) {
		e.PUT("/todos/:id", todoContoller.Edit, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		todoBody, _ := json.Marshal(map[string]interface{}{
			"name": "Healing",
			"todo_start": "2018-09-22T12:42:31+07:00",
			"todo_end": "2018-09-23T15:42:31+07:00",
		})

		req := httptest.NewRequest(echo.PUT, "/todos/1", bytes.NewBuffer(todoBody))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		var response common.ResponseSuccess

		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "success", response.Message)
		assert.Equal(t, "Healing", response.Data.(map[string]interface{})["name"])
		assert.Equal(t, "2018-09-22T12:42:31+07:00", response.Data.(map[string]interface{})["todo_start"])
		assert.Equal(t, "2018-09-23T15:42:31+07:00", response.Data.(map[string]interface{})["todo_end"])
		assert.Equal(t, "uncomplete", response.Data.(map[string]interface{})["status"])
	})
	
	t.Run("Update Todo Failed Not Found", func(t *testing.T) {
		e.PUT("/todos/:id", todoContoller.Edit, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		todoBody, _ := json.Marshal(map[string]interface{}{
			"name": "Healing",
			"todo_start": "2018-09-22T12:42:31+07:00",
			"todo_end": "2018-09-23T15:42:31+07:00",
		})

		wrongToken, _ := middlewares.CreateToken(3)

		req := httptest.NewRequest(echo.PUT, "/todos/1", bytes.NewBuffer(todoBody))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", wrongToken))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.NotNil(t, rec.Body)
	})
	
	t.Run("Update Todo Failed Bad Request", func(t *testing.T) {
		e.PUT("/todos/:id", todoContoller.Edit, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		todoBody, _ := json.Marshal(map[string]interface{}{
			"name": "Healing",
			"todo_start": "kjahdhka",
			"todo_end": "2018-09-23T15:42:31+07:00",
		})

		req := httptest.NewRequest(echo.PUT, "/todos/:id", bytes.NewBuffer(todoBody))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.NotNil(t, rec.Body)
	})
	
	t.Run("Update Todo Failed Unauthorize", func(t *testing.T) {
		e.PUT("/todos/:id", todoContoller.Edit, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		todoBody, _ := json.Marshal(map[string]interface{}{
			"name": "Healing",
			"todo_start": "2018-09-22T12:42:31+07:00",
			"todo_end": "2018-09-23T15:42:31+07:00",
		})

		req := httptest.NewRequest(echo.PUT, "/todos/1", bytes.NewBuffer(todoBody))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token + "wrongtoken"))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.NotNil(t, rec.Body)
	})
}

func TestDeleteTodo(t *testing.T)  {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	todoRepo := todoRepo.NewTodoRepo(db)
	todoContoller := todoContoller.NewTodoController(todoRepo)
	
	e := echo.New()

	t.Run("Delete Todo Success", func(t *testing.T) {
		e.DELETE("/todos/:id", todoContoller.Delete, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		req := httptest.NewRequest(echo.DELETE, "/todos/1", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		var response common.ResponseSuccess

		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "success", response.Message)
	})
	
	t.Run("Delete Todo Failed Bad Req", func(t *testing.T) {
		e.DELETE("/todos/:id", todoContoller.Delete, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		wrongToken, _ := middlewares.CreateToken(2)

		req := httptest.NewRequest(echo.DELETE, "/todos/dsd", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", wrongToken))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.NotNil(t, rec.Body)
	})

	t.Run("Delete Todo Failed Unauthorize", func(t *testing.T) {
		e.DELETE("/todos/:id", todoContoller.Delete, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		req := httptest.NewRequest(echo.DELETE, "/todos/7", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token + "wrongtoken"))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.NotNil(t, rec.Body)
	})
}

func TestSetCompleteTodo(t *testing.T)  {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	todoRepo := todoRepo.NewTodoRepo(db)
	todoContoller := todoContoller.NewTodoController(todoRepo)
	
	e := echo.New()
	
	t.Run("Set Complete Todo Success", func(t *testing.T) {
		e.PUT("/todos/:id/complete", todoContoller.SetComplete, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		req := httptest.NewRequest(echo.PUT, "/todos/1/complete", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		var response common.ResponseSuccess

		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "success", response.Message)
		assert.Equal(t, "complete", response.Data.(map[string]interface{})["status"])
	})
	
	t.Run("Set Complete Todo Failed Not Found", func(t *testing.T) {
		e.POST("/todos/:id/complete", todoContoller.SetComplete, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		wrongToken, _ := middlewares.CreateToken(3)

		req := httptest.NewRequest(echo.POST, "/todos/6/complete", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", wrongToken))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.NotNil(t, rec.Body)
	})
	
	t.Run("Set Complete Todo Failed Bad Request", func(t *testing.T) {
		e.POST("/todos/:id/complete", todoContoller.SetComplete, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		req := httptest.NewRequest(echo.POST, "/todos/:id/complete", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.NotNil(t, rec.Body)
	})
	
	t.Run("Set Complete Todo Failed Unauthorize", func(t *testing.T) {
		e.PUT("/todos/:id/complete", todoContoller.Edit, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		req := httptest.NewRequest(echo.PUT, "/todos/1/complete", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token + "wrongtoken"))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.NotNil(t, rec.Body)
	})
}

func TestReopenTodo(t *testing.T)  {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	todoRepo := todoRepo.NewTodoRepo(db)
	todoContoller := todoContoller.NewTodoController(todoRepo)
	
	e := echo.New()
	
	t.Run("Reopen Todo Success", func(t *testing.T) {
		e.POST("/todos/:id/reopen", todoContoller.Reopen, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		req := httptest.NewRequest(echo.POST, "/todos/1/reopen", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		var response common.ResponseSuccess

		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "uncomplete", response.Data.(map[string]interface{})["status"])
	})
	
	t.Run("Reopen Failed Not Found", func(t *testing.T) {
		e.POST("/todos/:id/reopen", todoContoller.SetComplete, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		wrongToken, _ := middlewares.CreateToken(3)

		req := httptest.NewRequest(echo.POST, "/todos/6/reopen", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", wrongToken))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.NotNil(t, rec.Body)
	})
	
	t.Run("Reopen Todo Failed Bad Request", func(t *testing.T) {
		e.POST("/todos/:id/reopen", todoContoller.SetComplete, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		req := httptest.NewRequest(echo.POST, "/todos/:id/reopen", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.NotNil(t, rec.Body)
	})
	
	t.Run("Reopen Todo Failed Unauthorize", func(t *testing.T) {
		e.PUT("/todos/:id/reopen", todoContoller.Edit, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		req := httptest.NewRequest(echo.PUT, "/todos/1/reopen", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token + "wrongtoken"))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.NotNil(t, rec.Body)
	})
}
