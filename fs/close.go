package fs

import "errors"

func Close() error {
	if fsFile == nil {
		return errors.New("Unable to close file. File is not even opened.")
	}

	if fsFile != nil {
		if err := fsFile.Close(); err != nil {
			return err
		}
	}

	return nil
}
