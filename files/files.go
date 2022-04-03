package files

import (
	"encoding/base64"
	"errors"
	"os"
)

func SaveFile(file string, path string, overwirite bool) error {
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

	} else if overwirite {
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
