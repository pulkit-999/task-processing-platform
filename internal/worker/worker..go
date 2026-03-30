package worker

import (
	"encoding/json"
	"fmt"
	"time"

	"job-queue-go/internal/metrics"
	"job-queue-go/internal/models"
	"job-queue-go/internal/queue"
)

type WorkerPool struct {
	Queue       *queue.RedisQueue
	WorkerCount int
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.WorkerCount; i++ {
		go wp.worker(i)
	}
	select {} // block forever
}

func (wp *WorkerPool) worker(id int) {
	queues := []string{"high", "medium", "low"}

	for {
		data, err := wp.Queue.DequeuePriority(queues)
		if err != nil {
			continue
		}

		var job models.Job
		json.Unmarshal([]byte(data), &job)

		fmt.Printf("🚀 Worker %d picked job %s (attempt %d)\n", id, job.ID, job.Attempts+1)

		err = process(job)

		if err != nil {
			job.Attempts++

			// Retry logic
			if job.Attempts >= job.MaxRetry {
				fmt.Printf("❌ Job %s moved to DLQ after %d attempts\n", job.ID, job.Attempts)

				wp.Queue.Enqueue("dlq", job)
				metrics.IncFailed()
				continue
			}

			backoff := time.Duration(job.Attempts*2) * time.Second
			fmt.Printf("🔁 Retrying job %s (attempt %d) after %v\n", job.ID, job.Attempts, backoff)

			time.Sleep(backoff)
			wp.Queue.Enqueue(job.Priority, job)
			continue
		}

		fmt.Printf("✅ Job %s completed successfully\n", job.ID)
		metrics.IncProcessed()
	}
}

func process(job models.Job) error {
	// Simulate failure for first 2 attempts
	if job.Attempts < 2 {
		return fmt.Errorf("temporary failure")
	}

	time.Sleep(1 * time.Second)
	return nil
}
