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
	projectContoller "github.com/furqonzt99/todo-api/delivery/controllers/project"
	userContoller "github.com/furqonzt99/todo-api/delivery/controllers/user"
	"github.com/furqonzt99/todo-api/delivery/middlewares"
	"github.com/furqonzt99/todo-api/models"
	projectRepo "github.com/furqonzt99/todo-api/repository/project"
	userRepo "github.com/furqonzt99/todo-api/repository/user"
	"github.com/furqonzt99/todo-api/utils"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M)  {
	config := configs.GetConfig()

	db := utils.InitDB(config)
	db.Migrator().DropTable(&models.User{})
	db.Migrator().DropTable(&models.Project{})
	db.Migrator().DropTable(&models.Project{})
	utils.InitialMigrate(db)

	userRepo := userRepo.NewUserRepo(db)
	userContoller := userContoller.NewUserController(userRepo)
	
	e := echo.New()

	e.Validator = &common.CustomValidator{Validator: validator.New()}

	e.POST("/register", userContoller.Register)

	registerBody, _ := json.Marshal(map[string]interface{}{
		"name": "Arif",
		"email": "arif@gmail.com",
		"password": "1234qwer",
	})

	req := httptest.NewRequest(echo.POST, "/register", bytes.NewBuffer(registerBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	e.POST("/login", userContoller.Login)

	loginBody, _ := json.Marshal(map[string]interface{}{
		"email": "arif@gmail.com",
		"password": "1234qwer",
	})

	reqLogin := httptest.NewRequest(echo.POST, "/login", bytes.NewBuffer(loginBody))
	reqLogin.Header.Set("Content-Type", "application/json")
	recLogin := httptest.NewRecorder()

	e.ServeHTTP(recLogin, reqLogin)

	var response common.ResponseSuccess

	json.Unmarshal(recLogin.Body.Bytes(), &response)

	fmt.Println(response)

	token = response.Data.(string)
	
	fmt.Println("TEST MAIN JALAN NIH")

	m.Run()

}

func TestInsertProject(t *testing.T)  {
	config := configs.GetConfig()

	db := utils.InitDB(config)
	db.Migrator().DropTable(&models.Project{})
	utils.InitialMigrate(db)

	projectRepo := projectRepo.NewProjectRepo(db)
	projectContoller := projectContoller.NewProjectController(projectRepo)
	
	e := echo.New()
	
	t.Run("Insert Project Success", func(t *testing.T) {
		e.POST("/projects", projectContoller.Insert, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		projectBody, _ := json.Marshal(map[string]interface{}{
			"name": "Metaverse",
			"description": "ini description",
		})

		req := httptest.NewRequest(echo.POST, "/projects", bytes.NewBuffer(projectBody))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		var response common.ResponseSuccess

		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "success", response.Message)
		assert.Equal(t, "Metaverse", response.Data.(map[string]interface{})["name"])
		assert.Equal(t, "ini description", response.Data.(map[string]interface{})["description"])
	})
	
	t.Run("Insert Project Failed Unauthorize", func(t *testing.T) {
		e.POST("/projects", projectContoller.Insert, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		projectBody, _ := json.Marshal(map[string]interface{}{
			"name": "Metaverse",
			"description": "ini description",
		})

		req := httptest.NewRequest(echo.POST, "/projects", bytes.NewBuffer(projectBody))
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

func TestGetAllProject(t *testing.T)  {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	projectRepo := projectRepo.NewProjectRepo(db)
	projectContoller := projectContoller.NewProjectController(projectRepo)
	
	e := echo.New()
	
	t.Run("Get All Project Success", func(t *testing.T) {
		e.GET("/projects", projectContoller.GetAll, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		req := httptest.NewRequest(echo.GET, "/projects", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		var response common.ResponseSuccess

		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.NotNil(t, response.Data)
	})
	
	t.Run("Get Project Failed Unauthorize", func(t *testing.T) {
		e.GET("/projects", projectContoller.GetAll, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		wrongToken, _ := middlewares.CreateToken(4)

		req := httptest.NewRequest(echo.GET, "/projects", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", wrongToken + "wrongtoken"))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.NotNil(t, rec.Body)
	})
}

func TestGetProject(t *testing.T)  {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	projectRepo := projectRepo.NewProjectRepo(db)
	projectContoller := projectContoller.NewProjectController(projectRepo)
	
	e := echo.New()
	
	t.Run("Get Project Success", func(t *testing.T) {
		e.GET("/projects/:id", projectContoller.Get, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		req := httptest.NewRequest(echo.GET, "/projects/1", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		var response common.ResponseSuccess

		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "Metaverse", response.Data.(map[string]interface{})["name"])
		assert.Equal(t, "ini description", response.Data.(map[string]interface{})["description"])
	})
	
	t.Run("Get Project Failed Not Found", func(t *testing.T) {
		e.GET("/projects/:id", projectContoller.Get, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		wrongToken, _ := middlewares.CreateToken(4)

		req := httptest.NewRequest(echo.GET, "/projects/4", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", wrongToken))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.NotNil(t, rec.Body)
	})
	
	t.Run("Get Project Failed Unauthorize", func(t *testing.T) {
		e.GET("/projects/:id", projectContoller.Get, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		wrongToken, _ := middlewares.CreateToken(4)

		req := httptest.NewRequest(echo.GET, "/projects/1", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", wrongToken + "wrongtoken"))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.NotNil(t, rec.Body)
	})
}

func TestUpdateProject(t *testing.T)  {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	projectRepo := projectRepo.NewProjectRepo(db)
	projectContoller := projectContoller.NewProjectController(projectRepo)
	
	e := echo.New()
	
	t.Run("Update Project Success", func(t *testing.T) {
		e.PUT("/projects/:id", projectContoller.Edit, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		projectBody, _ := json.Marshal(map[string]interface{}{
			"name": "Metaverse Nih",
			"description": "ini description",
		})

		req := httptest.NewRequest(echo.PUT, "/projects/1", bytes.NewBuffer(projectBody))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		var response common.ResponseSuccess

		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "success", response.Message)
		assert.Equal(t, "Metaverse Nih", response.Data.(map[string]interface{})["name"])
		assert.Equal(t, "ini description", response.Data.(map[string]interface{})["description"])
	})
	
	t.Run("Update Project Failed Not Found", func(t *testing.T) {
		e.PUT("/projects/:id", projectContoller.Edit, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		projectBody, _ := json.Marshal(map[string]interface{}{
			"name": "Metaverse",
			"description": "ini description",
		})

		wrongToken, _ := middlewares.CreateToken(3)

		req := httptest.NewRequest(echo.PUT, "/projects/1", bytes.NewBuffer(projectBody))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", wrongToken))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.NotNil(t, rec.Body)
	})
	
	t.Run("Update Project Failed Bad Request", func(t *testing.T) {
		e.PUT("/projects/:id", projectContoller.Edit, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		projectBody, _ := json.Marshal(map[string]interface{}{
			"name": "Metaverse",
			"description": "ini description",
		})

		req := httptest.NewRequest(echo.PUT, "/projects/:id", bytes.NewBuffer(projectBody))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.NotNil(t, rec.Body)
	})
	
	t.Run("Update Project Failed Unauthorize", func(t *testing.T) {
		e.PUT("/projects/:id", projectContoller.Edit, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		projectBody, _ := json.Marshal(map[string]interface{}{
			"name": "Metaverse",
			"description": "ini description",
		})

		req := httptest.NewRequest(echo.PUT, "/projects/1", bytes.NewBuffer(projectBody))
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token + "wrongtoken"))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.NotNil(t, rec.Body)
	})
}

func TestDeleteProject(t *testing.T)  {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	projectRepo := projectRepo.NewProjectRepo(db)
	projectContoller := projectContoller.NewProjectController(projectRepo)
	
	e := echo.New()

	t.Run("Delete Project Success", func(t *testing.T) {
		e.DELETE("/projects/:id", projectContoller.Delete, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		req := httptest.NewRequest(echo.DELETE, "/projects/1", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		var response common.ResponseSuccess

		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "success", response.Message)
	})
	
	t.Run("Delete Project Failed Bad Req", func(t *testing.T) {
		e.DELETE("/projects/:id", projectContoller.Delete, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		wrongToken, _ := middlewares.CreateToken(2)

		req := httptest.NewRequest(echo.DELETE, "/projects/dsd", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", wrongToken))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.NotNil(t, rec.Body)
	})

	t.Run("Delete Project Failed Unauthorize", func(t *testing.T) {
		e.DELETE("/projects/:id", projectContoller.Delete, middleware.JWT([]byte(constants.JWT_SECRET_KEY)))

		req := httptest.NewRequest(echo.DELETE, "/projects/:id", nil)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token + "wrongtoken"))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		assert.NotNil(t, rec.Body)
	})
}