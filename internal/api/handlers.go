package api

import (
	"digantara/internal/db"
	"digantara/internal/scheduler"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
)

var validate = validator.New()

func allJobs(c *gin.Context) {
	jobs, err := db.GetAllJobs()
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Error fetching jobs",
			"error":   "error getting list of jobs",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "All jobs",
		"jobs":    jobs,
	})
}

func createJob(c *gin.Context) {
	var req CreateJobRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := scheduler.AddJob(req.Type, req.Cron, req.Message, req.Name)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "job created",
		"id":      id,
		"req":     req,
	})
}

func getJobByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "invalid job id",
			"id":    id,
		})
		return
	}

	job, err := db.GetJobByID(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "job does not exist",
			"id":    id,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Job Details",
		"job":     job,
	})
}

// todo: func to evaluate high level data like Monday, Everyday and generate cron expression instead of sending cron expression directly in the request data
func JobExpression() {

	fmt.Println("::::")
}
