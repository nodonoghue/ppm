package generate

import (
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/nodonoghue/ppm/internal/models"
)

func Password(ch chan<- string, wg *sync.WaitGroup, configuration models.CommandFlags) {
	defer wg.Done()
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// //use shuffle to create a len20 password and print to terminal:
	// buf := make([]byte, models.Length)
	// for i := 0; i < models.Length; i++ {
	// 	buf[i] = models.AllChars[rand.Intn(len(models.AllChars))]
	// }

	var builder strings.Builder

	if *configuration.NumUpperCase > 0 {
		builder.Write([]byte(getUpperChars(*configuration.NumUpperCase)))
	}
	if *configuration.NumNumbers > 0 {
		builder.Write([]byte(getNumberChars(*configuration.NumNumbers)))
	}
	if *configuration.NumSpecial > 0 {
		builder.Write([]byte(getSpecialChars(*configuration.NumSpecial)))
	}

	numLower := getNumLower(configuration)
	if numLower > 0 {
		builder.Write([]byte(getLowerChars(numLower)))
	}

	ch <- shuffleChars(strings.Split(builder.String(), ""))
}

func getUpperChars(numUpper int) string {
	buf := make([]byte, numUpper)
	for i := 0; i < numUpper; i++ {
		buf[i] = models.UpperCase[rand.Intn(len(models.UpperCase))]
	}
	return string(buf)
}

func getNumberChars(numNumbers int) string {
	buf := make([]byte, numNumbers)
	for i := 0; i < numNumbers; i++ {
		buf[i] = models.Numbers[rand.Intn(len(models.Numbers))]
	}
	return string(buf)
}

func getSpecialChars(numSpecialChars int) string {
	buf := make([]byte, numSpecialChars)
	for i := 0; i < numSpecialChars; i++ {
		buf[i] = models.SpecialChars[rand.Intn(len(models.SpecialChars))]
	}
	return string(buf)
}

func getLowerChars(numLower int) string {
	buf := make([]byte, numLower)
	for i := 0; i < numLower; i++ {
		buf[i] = models.LowerCase[rand.Intn(len(models.LowerCase))]
	}
	return string(buf)
}

func shuffleChars(chars []string) string {
	rand.Shuffle(len(chars), func(i, j int) {
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
