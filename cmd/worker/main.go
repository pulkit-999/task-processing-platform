package main

import (
	"job-queue-go/internal/queue"
	"job-queue-go/internal/worker"
)

func main() {
	q := queue.NewRedisQueue()

	wp := worker.WorkerPool{
		Queue:       q,
		WorkerCount: 5,
	}

	wp.Start()
}
