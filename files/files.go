package files

import (
	"encoding/base64"
	"errors"
	"os"
)

// as I just cannot find a method to use golang to read the files in the .app/Contents/Resources, I stored necessary files as base64 encoding in the files subpackage

func SaveFile(file string, path string, overwrite bool) error {
	// save a base64 encoded file to the specified path
	// file is the base64 encoded data of a file
	// if there already exists a file at path, true value of overwrite will replace that file

	f, err := base64.StdEncoding.DecodeString(file)
	if err != nil {
		return err
	}

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		f1, err := os.Create(path)
		if err != nil {
			return err
		}
		defer f1.Close()

		_, err = f1.Write(f)
		if err != nil {
			return err
		}

	} else if overwrite {
		f1, err := os.Create(path)
		if err != nil {
			return err
		}
		defer f1.Close()

		_, err = f1.Write(f)
		if err != nil {
			return err
		}
	}

	return nil

}
