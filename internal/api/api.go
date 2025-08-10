package api

import (
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/time", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"time": time.Now(),
		})
	})

	jobs := r.Group("/jobs")
	{
		jobs.GET("", allJobs)
		jobs.POST("", createJob)
		jobs.GET("/id/:id", getJobByID)
	}

	// r.GET("/jobs", AllJobs)

	return r
}
