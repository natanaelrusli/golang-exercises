package downloader_concurrency

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/natanaelrusli/image-downloader/internal/directory"
)

const URL = "https://picsum.photos/200/300"
const ImageDirectory = "./images"

func ExecuteDownload(qty int, baseURL, baseFileName string) error {
	err := directory.CreateDirectory(ImageDirectory)

	if err != nil {
		return fmt.Errorf("error creating directory: %v", err.Error())
	}

	pool := NewWorkerPool(10) // Set the number of workers as needed
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
