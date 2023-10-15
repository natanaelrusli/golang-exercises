package downloader_concurrency

import "sync"

type workerPool struct {
	maxWorker  int
	queuedTask chan func()
	wg         sync.WaitGroup
}

// NewWorkerPool creates a new worker pool with the specified number of workers.
func NewWorkerPool(maxWorker int) *workerPool {
	return &workerPool{
		maxWorker:  maxWorker,
		queuedTask: make(chan func()),
	}
}

// Submit submits a task to the worker pool.
func (wp *workerPool) Submit(task func()) {
	wp.wg.Add(1)
	wp.queuedTask <- task
}

// Wait waits for all tasks to be completed.
func (wp *workerPool) Wait() {
	close(wp.queuedTask)
	wp.wg.Wait()
}

// worker runs tasks from the queuedTask channel.
func (wp *workerPool) worker() {
	for task := range wp.queuedTask {
		task()
		wp.wg.Done()
	}
}

// StartWorkers starts worker goroutines.
func (wp *workerPool) StartWorkers() {
	for i := 0; i < wp.maxWorker; i++ {
		go wp.worker()
	}
}
