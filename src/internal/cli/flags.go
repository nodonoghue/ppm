package cli

import (
	"flag"

	"github.com/nodonoghue/ppm/internal/models"
)

func GetFlags() models.CommandFlags {
	var commandFlags models.CommandFlags
	commandFlags.Options = setNumOptsLong()
	flag.Parse()
	return commandFlags
}

func setNumOptsLong() *int {
	return flag.Int("o", 1, "Sets the number of password options to generate")
}
