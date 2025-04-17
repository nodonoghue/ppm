package cli

import (
	"flag"
	"os"

	"github.com/nodonoghue/ppm/internal/models"
)

var exitFunc = os.Exit

func GetFlags() models.CommandFlags {
	var commandFlags models.CommandFlags

	commandFlags.NumVariants = setNumVariants()
	commandFlags.Length = setPasswordLength()
	commandFlags.NumUpperCase = setNumUpperCase()
	commandFlags.NumNumbers = setNumNumbers()
	commandFlags.NumSpecial = setNumSpecial()

	showHelp := helpFlag()

	if showHelp {
		flag.Usage()
		exitFunc(0)
	}

	flag.Parse()
	return commandFlags
}

func setNumVariants() *int {
	val := flag.Int("v", 0, "Sets the number of password variants to generate")
	if *val == 0 {
		return val
	}
	return val
}

func setPasswordLength() *int {
	val := flag.Int("l", 0, "Sets the password length, must be at least 8 chars long")
	if *val == 0 {
		*val = 8
		return val
	}
	return val
}

func setNumUpperCase() *int {
	val := flag.Int("u", 0, "Sets the number of upper case chars in each variant")
	if *val == 0 {
		return val
	}
	return val
}

func setNumNumbers() *int {
	val := flag.Int("n", 0, "Sets the number of number chars in each variant")
	if *val == 0 {
		return val
	}
	return val
}

func setNumSpecial() *int {
	val := flag.Int("s", 0, "Sets the number of special ( ! @ # $ % ^ & * ) chars in each variant")
	if *val == 0 {
		return val
	}
	return val
}

func helpFlag() bool {
	val := flag.Bool("h", false, "Prints this help message")
	return *val
}
