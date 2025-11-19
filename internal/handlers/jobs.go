package handlers

import (
	"net/http"

	"github.com/Ademayowa/learn-d-compose/internal/models"

	"github.com/gin-gonic/gin"
)

// Create a job
func createJob(ctx *gin.Context) {
	var job models.Job

	err := ctx.ShouldBindJSON(&job)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "could not parse job data: " + err.Error()})
		return
	}

	err = job.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not save job: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "job created", "job": job})
}

// Get all jobs
func getJobs(ctx *gin.Context) {
	jobs, err := models.GetAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch jobs"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"jobs": jobs})
}
