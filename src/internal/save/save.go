package save

import (
	"os"

	"github.com/nodonoghue/ppm/internal/models"
)

// write a file, simple plain text output for now, will need to think of a system to
// store the file based on some initial user input and configuration, then create a
// pass phrase system to encrypt and decrypt the file for safe storage of the passwords
func SaveValue(val string) error {
	if !checkFile("") {
		var error models.GeneralError
		error.FunctionName = "save.checkFile"
		error.Message = "File does not exist"
		return error
	}
	return nil
}

func checkFile(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		return false
	}
	return true
}

func writeFile() {

}
