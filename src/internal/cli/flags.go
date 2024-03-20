package cli

import (
	"flag"

	"github.com/nodonoghue/ppm/internal/models"
)

func GetFlags() models.CommandFlags {
	var commandFlags models.CommandFlags
	commandFlags.NumVariants = setNumVariants()
	commandFlags.Length = setPasswordLength()
	commandFlags.NumLowerCase = setNumLowerCase()
	commandFlags.NumUpperCase = setNumUpperCase()
	commandFlags.NumNumbers = setNumNumbers()
	commandFlags.NumSpecial = setNumSpecial()
	flag.Parse()
	return commandFlags
}

func setNumVariants() *int {
	return flag.Int("v", 1, "Sets the number of password variants to generate")
}

func setPasswordLength() *int {
	return flag.Int("length", 8, "Sets the password length, must be at least 8 chars long")
}

func setNumUpperCase() *int {
	return flag.Int("u", 1, "Sets the number of upper case chars in each variant")
}

func setNumLowerCase() *int {
	return flag.Int("l", 1, "Sets the number of lower case chars in each variant.  Will be overwritten if the sum of all char settings is less than 8 to ensure the password length is at least 8 chars.")
}

func setNumNumbers() *int {
	return flag.Int("n", 1, "Sets the number of number chars in each variant")
}

func setNumSpecial() *int {
	return flag.Int("s", 0, "Sets the number of special ( ! @ # $ % ^ & * ) chars in each variant")
}
