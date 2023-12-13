package fs

import (
	"os"
)

var fsFile *os.File

func Open(f string) error {
	file, err := os.Open(f)

	if err != nil {
		return err
	}

	fsFile = file

	return nil
}
