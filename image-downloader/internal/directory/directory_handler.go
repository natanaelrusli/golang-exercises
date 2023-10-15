package directory

import (
	"os"
)

func CreateDirectory(directory string) error {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err := os.Mkdir(directory, os.ModePerm)

		if err != nil {
			return err
		}
	}

	return nil
}
