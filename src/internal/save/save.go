package save

import (
	"os"

	"github.com/nodonoghue/ppm/internal/models/constants"
)

func SaveValue(val string) error {
	return writeFile(val)
}

func openFile(filename string) (*os.File, error) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func writeFile(val string) error {
	filename := constants.Filename
	file, err := openFile(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(val + "\n"); err != nil {
		return err
	}
	return nil
}
