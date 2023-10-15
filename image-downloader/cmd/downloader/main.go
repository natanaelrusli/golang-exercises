package main

import (
	"log"

	"github.com/natanaelrusli/image-downloader/internal/downloader"
)

const URL = "https://picsum.photos/200/300"
const ImageDirectory = "./images"

func main() {
	fileName := "sample"
	err := downloader.ExecuteDownload(100, URL, fileName)

	if err != nil {
		log.Fatal(err.Error())
	}
}
