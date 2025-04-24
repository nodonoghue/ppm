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
		log.Fatal("Unable to read input", err)
	}

	return strings.Replace(input, "\n", "", -1)
}

func GetVariant(index string) (string, error) {
	num, err := strconv.Atoi(index)
	if err != nil {
		return "", err
	}
	return models.Variants[num-1], nil
}
