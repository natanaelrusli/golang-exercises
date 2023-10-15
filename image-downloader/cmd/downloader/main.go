package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const URL = "https://picsum.photos/200/300"
const ImageDirectory = "./images"

func main() {
	fileName := "sample"
	err := ExecuteDownload(100, URL, fileName)

	if err != nil {
		log.Fatal(err.Error())
	}
}

func ExecuteDownload(qty int, baseURL, baseFileName string) error {
	err := createDirectory(ImageDirectory)

	if err != nil {
		return fmt.Errorf("error creating directory: %v", err.Error())
	}

	for i := 1; i <= qty; i++ {
		err := downloadFile(URL, fmt.Sprintf("%s %v.png", baseFileName, i))

		if err != nil {
			return fmt.Errorf("error downloading from: %s", URL)
		}
	}

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

	fmt.Printf("File %s downloaded in ./images direcrory\n", fileName)

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
