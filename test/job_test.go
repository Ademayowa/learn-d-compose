package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	db "github.com/Ademayowa/learn-d-compose/internal/database"
	"github.com/Ademayowa/learn-d-compose/internal/handlers"
	"github.com/Ademayowa/learn-d-compose/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func setupTest(t *testing.T) *gin.Engine {
	gin.SetMode(gin.TestMode)

	// Load .env.test only in local development (ignored in CI)
	godotenv.Load("../.env.test")

	db.InitDB()

	// Clean up jobs table before each test
	_, err := db.DB.Exec("DELETE FROM jobs")
	assert.NoError(t, err)

	router := gin.New()
	handlers.RegisterRoutes(router)
	return router
}

func TestCreateJob(t *testing.T) {
	router := setupTest(t)

	job := models.Job{
		Title:       "Backend Engineer",
		Description: "Work with Go and PostgreSQL",
	}

	body, _ := json.Marshal(job)
	req := httptest.NewRequest(http.MethodPost, "/jobs", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response models.Job
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.NotEmpty(t, response.ID)
	assert.Equal(t, job.Title, response.Title)
}

func TestGetJobs(t *testing.T) {
	router := setupTest(t)

	// Create test jobs
	jobs := []models.Job{
		{Title: "Frontend Developer", Description: "React expert"},
		{Title: "DevOps Engineer", Description: "AWS and Docker"},
	}

	for _, job := range jobs {
		job.Save()
	}

	req := httptest.NewRequest(http.MethodGet, "/jobs", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []models.Job
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Len(t, response, 2)
}

func TestCreateJobValidation(t *testing.T) {
	router := setupTest(t)

	invalidJob := map[string]string{"title": ""}
	body, _ := json.Marshal(invalidJob)

	req := httptest.NewRequest(http.MethodPost, "/jobs", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
