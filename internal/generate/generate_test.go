package generate

import (
	"math/rand"
	"testing"
	"time"
)

const (
	upperCaseCompare    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerCaseCompare    = "abcdefghijklmnopqrstuvwxyz"
	numberCompare       = "0123456789"
	specialCharsCompare = "!@#$%^&*"

	lengthError string = "expected length %d, but got %d"
	charError   string = "Character '%c' is not in the %s set"
)

func TestGetUpperChars(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	numUpper := 5
	result := getUpperChars(numUpper, r)

	if len(result) != numUpper {
		t.Errorf(lengthError, numUpper, len(result))
	}

	for _, char := range result {
		if !contains(upperCaseCompare, char) {
			t.Errorf(charError, char, "upperCaseCompare")
		}
	}
}

func TestGetZeroUpper(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	numUpper := 0
	result := getUpperChars(numUpper, r)

	if len(result) != numUpper {
		t.Errorf(lengthError, numUpper, len(result))
	}
}

func TestGetLowerCase(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	numLower := 10
	result := getLowerChars(numLower, r)

	if len(result) != numLower {
		t.Errorf(lengthError, numLower, len(result))
	}

	for _, char := range result {
		if !contains(lowerCaseCompare, char) {
			t.Errorf(charError, char, "lowerCaseCompare")
		}
	}
}

func TestGetZeroLower(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	numLower := 0
	result := getLowerChars(numLower, r)

	if len(result) != numLower {
		t.Errorf(lengthError, numLower, len(result))
	}
}

func TestGetSpecialChars(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	numSpecial := 4
	result := getSpecialChars(numSpecial, r)

	if len(result) != numSpecial {
		t.Errorf(lengthError, numSpecial, len(result))
	}

	for _, char := range result {
		if !contains(specialCharsCompare, char) {
			t.Errorf(charError, char, "specialCharsCompare")
		}
	}
}

func TestGetZeroSpecial(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	numSpecial := 0
	result := getSpecialChars(numSpecial, r)

	if len(result) != numSpecial {
		t.Errorf(lengthError, numSpecial, len(result))
	}
}

func TestGetNumbers(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	numNumbers := 6
	result := getNumberChars(numNumbers, r)

	if len(result) != numNumbers {
		t.Errorf(lengthError, numNumbers, len(result))
	}

	for _, char := range result {
		if !contains(numberCompare, char) {
			t.Errorf(charError, char, "numberCompare")
		}
	}
}

func TestGetZeroNumbers(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	numNumbers := 0
	result := getNumberChars(numNumbers, r)

	if len(result) != numNumbers {
		t.Errorf(lengthError, numNumbers, len(result))
	}
}

// Helper function to check if a character exists in a string
func contains(s string, c rune) bool {
	for _, r := range s {
		if r == c {
			return true
		}
	}
	return false
}
