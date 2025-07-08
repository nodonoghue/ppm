package save

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/nodonoghue/ppm/internal/models/constants"
)

func OverwriteFile(val []byte) error {
	return writeFile(val)
}

func writeFile(val []byte) error {
	filename := constants.Filename
	file, err := openFile(filename)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	if _, err := file.Write(val); err != nil {
		return err
	}
	return nil
}

func openFile(filename string) (*os.File, error) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func ReadFile() ([]byte, error) {
	filename := constants.Filename
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return nil, nil
	}
	return ioutil.ReadFile(filename)
}