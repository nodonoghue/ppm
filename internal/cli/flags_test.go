package cli

import (
	"os"
	"testing"
)

func mockExit(code int) {
	panic("exit called")
}

func TestFlags(t *testing.T) {
	os.Args = []string{"cmd", "-v=5", "-l=15", "-u=3", "-n=3", "-s=2"}

	flags := GetFlags()

	if *flags.NumVariants != 5 {
		t.Errorf("Number of variants should be 5, got %d", *flags.NumVariants)
	}

	if *flags.Length != 15 {
		t.Errorf("Length should be 15, got %d", *flags.Length)
	}

	if *flags.NumUpperCase != 3 {
		t.Errorf("NumUpperCase should be 3, got %d", *flags.NumUpperCase)
	}

	if *flags.NumNumbers != 3 {
		t.Errorf("NumNumbers should be 3, got %d", *flags.NumNumbers)
	}

	if *flags.NumSpecial != 2 {
		t.Errorf("NumSpecial should be 2, got %d", *flags.NumSpecial)
	}
}

func TestGetFlags_Help(t *testing.T) {
	originalExit := exitFunc
	defer func() {
		exitFunc = originalExit
	}()

	exitFunc = func(code int) { panic("exit called") }

	os.Args = []string{"cmd", "-h"}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected exitFunc to be called, but it wasn't")
		}
	}()

	GetFlags()
}

func TestGetFlags_InvalidLength(t *testing.T) {
	originalExit := exitFunc
	defer func() {
		exitFunc = originalExit
	}()

	exitFunc = func(code int) { panic("exit called") }

	os.Args = []string{"cmd", "-l=7"}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected exitFunc to be called, but it wasn't")
		}
	}()

	GetFlags()
}

func TestGetFlags_InvalidSum(t *testing.T) {
	originalExit := exitFunc
	defer func() {
		exitFunc = originalExit
	}()

	exitFunc = func(code int) { panic("exit called") }

	os.Args = []string{"cmd", "-l=10", "-u=5", "-n=5", "-s=1"}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected exitFunc to be called, but it wasn't")
		}
	}()

	GetFlags()
}
