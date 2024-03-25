package cli

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/nodonoghue/ppm/internal/models"
)

func ReadInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Unable to read input")
	}

	return strings.Replace(input, "\n", "", -1)
}

func GetVariant(index string) string {
	num, err := strconv.Atoi(index)
	if err != nil {
		log.Fatal("Input must be numeric. ", err.Error())
	}
	return models.Variants[num-1]
}
