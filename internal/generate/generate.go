package generate

import (
	"math/rand"
	"strings"
	"sync"

	"github.com/nodonoghue/ppm/internal/models"
	"github.com/nodonoghue/ppm/internal/models/constants"
)

func Password(ch chan<- string, wg *sync.WaitGroup, configuration models.CommandFlags, r *rand.Rand) {
	defer wg.Done()

	var builder strings.Builder
	builder.Grow(*configuration.Length)

	if *configuration.NumUpperCase > 0 {
		builder.Write([]byte(getUpperChars(*configuration.NumUpperCase, r)))
	}
	if *configuration.NumNumbers > 0 {
		builder.Write([]byte(getNumberChars(*configuration.NumNumbers, r)))
	}
	if *configuration.NumSpecial > 0 {
		builder.Write([]byte(getSpecialChars(*configuration.NumSpecial, r)))
	}

	numLower := getNumLower(configuration)
	if numLower > 0 {
		builder.Write([]byte(getLowerChars(numLower, r)))
	}

	ch <- shuffleChars(strings.Split(builder.String(), ""), r)
}

func getUpperChars(numUpper int, r *rand.Rand) string {
	buf := make([]byte, numUpper)
	for i := range numUpper {
		buf[i] = constants.UpperCase[r.Intn(len(constants.UpperCase))]
	}
	return string(buf)
}

func getNumberChars(numNumbers int, r *rand.Rand) string {
	buf := make([]byte, numNumbers)
	for i := range numNumbers {
		buf[i] = constants.Numbers[r.Intn(len(constants.Numbers))]
	}
	return string(buf)
}

func getSpecialChars(numSpecialChars int, r *rand.Rand) string {
	buf := make([]byte, numSpecialChars)
	for i := range numSpecialChars {
		buf[i] = constants.SpecialChars[r.Intn(len(constants.SpecialChars))]
	}
	return string(buf)
}

func getLowerChars(numLower int, r *rand.Rand) string {
	buf := make([]byte, numLower)
	for i := range numLower {
		buf[i] = constants.LowerCase[r.Intn(len(constants.LowerCase))]
	}
	return string(buf)
}

func shuffleChars(chars []string, r *rand.Rand) string {
	r.Shuffle(len(chars), func(i, j int) {
		chars[i], chars[j] = chars[j], chars[i]
	})
	return strings.Join(chars, "")
}

func getNumLower(configuration models.CommandFlags) int {
	//sum  up all parts to determine the leftover, this leftover will be the number of lowercase
	if (*configuration.NumUpperCase + *configuration.NumNumbers + *configuration.NumSpecial) < *configuration.Length {
		return *configuration.Length - (*configuration.NumUpperCase + *configuration.NumNumbers + *configuration.NumSpecial)
	}
	return 0
}
