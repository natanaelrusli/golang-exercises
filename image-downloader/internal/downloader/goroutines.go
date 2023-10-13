package downloader

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

type workerPool struct {
	maxWorker  int
	queuedTask chan func()
	wg         sync.WaitGroup
}

const URL = "https://picsum.photos/200/300"
const ImageDirectory = "./images"

func main() {
	fileName := "sample"
	err := ExecuteDownload(10, URL, fileName)

	if err != nil {
		log.Fatal(err.Error())
	}
}

func ExecuteDownload(qty int, baseURL, baseFileName string) error {
	err := createDirectory(ImageDirectory)

	if err != nil {
		return fmt.Errorf("error creating directory: %v", err.Error())
	}

	pool := NewWorkerPool(3) // Set the number of workers as needed
	pool.StartWorkers()

	for i := 1; i <= qty; i++ {
		index := i
		pool.Submit(func() {
			err := downloadFile(URL, fmt.Sprintf("%s %v.png", baseFileName, index))
			if err != nil {
				log.Printf("Error downloading from %s: %v", URL, err)
			}
		})
	}

	pool.Wait()
	return nil
}

func downloadFile(URL, fileName string) error {
	fmt.Printf("Downloading %s from %s\n", fileName, URL)

	res, err := http.Get(URL)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return errors.New("received non 200 response code")
	}

	file, err := os.Create(fmt.Sprintf("./images/%s", fileName))

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, res.Body)

	if err != nil {
		return err
	}

	fmt.Printf("File %s downloaded in ./images directory\n", fileName)

	return nil
}

func createDirectory(directory string) error {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err := os.Mkdir(directory, os.ModePerm)

		if err != nil {
			return err
		}
	}

	return nil
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
