package cli

import (
	"flag"
	"os"

	"github.com/nodonoghue/ppm/internal/models"
)

func GetFlags() models.CommandFlags {
	var commandFlags models.CommandFlags
	var isPresent bool
	commandFlags.NumVariants, isPresent = setNumVariants()
	commandFlags.Length, isPresent = setPasswordLength()
	commandFlags.NumLowerCase, isPresent = setNumLowerCase()
	commandFlags.NumUpperCase, isPresent = setNumUpperCase()
	commandFlags.NumNumbers, isPresent = setNumNumbers()
	commandFlags.NumSpecial, isPresent = setNumSpecial()

	if !isPresent {
		flag.Usage()
		os.Exit(0)
	}
	flag.Parse()
	return commandFlags
}

func setNumVariants() (*int, bool) {
	val := flag.Int("v", 0, "Sets the number of password variants to generate")
	if *val == 0 {
		return val, false
	}
	return val, true
}

func setPasswordLength() (*int, bool) {
	val := flag.Int("length", 0, "Sets the password length, must be at least 8 chars long")
	if *val == 0 {
		return val, false
	}
	return val, true
}

func setNumUpperCase() (*int, bool) {
	val := flag.Int("u", 0, "Sets the number of upper case chars in each variant")
	if *val == 0 {
		return val, false
	}
	return val, true
}

func setNumLowerCase() (*int, bool) {
	val := flag.Int("l", 0, "Sets the number of lower case chars in each variant.  Will be overwritten if the sum of all char settings is less than 8 to ensure the password length is at least 8 chars.")
	if *val == 0 {
		return val, false
	}
	return val, true
}

func setNumNumbers() (*int, bool) {
	val := flag.Int("n", 0, "Sets the number of number chars in each variant")
	if *val == 0 {
		return val, false
	}
	return val, true
}

func setNumSpecial() (*int, bool) {
	val := flag.Int("s", 0, "Sets the number of special ( ! @ # $ % ^ & * ) chars in each variant")
	if *val == 0 {
		return val, false
	}
	return val, true
}
