package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"job-queue-go/internal/metrics"
	"job-queue-go/internal/models"
	"job-queue-go/internal/queue"
)

func main() {
	r := gin.Default()
	q := queue.NewRedisQueue()

	r.POST("/job", func(c *gin.Context) {
		var job models.Job

		if err := c.BindJSON(&job); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		job.ID = uuid.New().String()
		job.CreatedAt = time.Now()
		job.MaxRetry = 3

		if job.Priority == "" {
			job.Priority = "low"
		}

		err := q.Enqueue(job.Priority, job)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		metrics.IncQueued()
		c.JSON(http.StatusOK, gin.H{
			"message": "job created",
			"job_id":  job.ID,
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "API running"})
	})

	r.GET("/metrics", func(c *gin.Context) {
		c.JSON(http.StatusOK, metrics.M)
	})

	r.Run(":8080")
}
