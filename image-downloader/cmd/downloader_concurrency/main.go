package main

import (
	"log"

	downloaderconcurrency "github.com/natanaelrusli/image-downloader/internal/downloader_concurrency"
)

const URL = "https://picsum.photos/200/300"
const ImageDirectory = "./images"

func main() {
	fileName := "sample"
	err := downloaderconcurrency.ExecuteDownload(100, URL, fileName)

	if err != nil {
		log.Fatal(err.Error())
	}
}
