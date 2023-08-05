package worker

import (
	"context"
	"testing"
	"time"

	"go.uber.org/goleak"
)

func TestRunWorker(t *testing.T) {
	defer goleak.VerifyNone(t)
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create a task
	task := &Task{ID: 1}

	// Run the worker
	queue := RunWorker(ctx, 1000)

	// Add the task to the queue
	queue <- task

	// Allow the worker enough time to process the task
	time.Sleep(5 * time.Second)

	// Check if the worker is done by trying to add another task.
	// If the worker is not done, this will block, causing the test to fail due to timeout
	select {
	case queue <- task:
		// Task was added to the queue, which means the worker is done
	case <-ctx.Done():
		// Timeout, worker is not done
		t.Fatal("Worker did not finish processing in time")
	}
}
