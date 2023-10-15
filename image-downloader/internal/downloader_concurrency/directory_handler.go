package downloader_concurrency

import (
	"os"
)

const URL = "https://picsum.photos/200/300"
const ImageDirectory = "./images"

func createDirectory(directory string) error {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err := os.Mkdir(directory, os.ModePerm)

		if err != nil {
			return err
		}
	}

	return nil
}
