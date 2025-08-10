package api

import (
	"digantara/internal/scheduler"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func allJobs(c *gin.Context) {
	c.JSON(200, gin.H{"message": "All jobs"})
}

func createJob(c *gin.Context) {
	var req CreateJobRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := scheduler.AddJob(req.Type, req.Cron, req.Message)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{
		"message": "job created",
		"id":      id,
		"req":     req,
	})
}

func getJob(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Job Details",
		"id":      "",
	})
}

// todo: func to evaluate high level data like Monday, Everyday and generate cron expression instead of sending cron expression directly in the request data
func JobExpression() {

	fmt.Println("::::")
}
