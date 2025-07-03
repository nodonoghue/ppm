package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/nodonoghue/ppm/internal/models"
)

var exitFunc = os.Exit

func GetFlags() models.CommandFlags {
	var commandFlags models.CommandFlags

	commandFlags.NumVariants = flag.Int("v", 5, "Sets the number of password variants to generate")
	commandFlags.Length = flag.Int("l", 12, "Sets the password length, must be at least 8 chars long")
	commandFlags.NumUpperCase = flag.Int("u", 2, "Sets the number of upper case chars in each variant")
	commandFlags.NumNumbers = flag.Int("n", 2, "Sets the number of number chars in each variant")
	commandFlags.NumSpecial = flag.Int("s", 2, "Sets the number of special ( ! @ # $ % ^ & * ) chars in each variant")
	showHelp := flag.Bool("h", false, "Prints this help message")

	flag.Parse()

	if *showHelp {
		flag.Usage()
		exitFunc(0)
	}

	if *commandFlags.Length < 8 {
		fmt.Println("Password length must be at least 8 characters")
		exitFunc(1)
	}

	if (*commandFlags.NumUpperCase + *commandFlags.NumNumbers + *commandFlags.NumSpecial) > *commandFlags.Length {
		fmt.Println("The sum of uppercase, number, and special characters cannot be greater than the password length")
		exitFunc(1)
	}

	return commandFlags
}
