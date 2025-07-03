package cli

import (
	"os"
	"testing"
)

func TestInput(t *testing.T) {
	// Backup the original os.Stdin, then defer restoring for future tests
	originalStdin := os.Stdin
	defer func() { os.Stdin = originalStdin }()

	// Create a temporary file to simulate stdin, then defer cleaning up temp file
	tempFile, err := os.CreateTemp("", "mock_stdin")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer func(name string) {
		_ = os.Remove(name)
	}(tempFile.Name())

	// Write mock input to the temp file
	want := "test"
	if _, err := tempFile.Write([]byte(want + "\n")); err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}

	// Rewind the file pointer to the beginning of the file to allow for reading
	// then set os.Stdin = tempFile
	if _, err := tempFile.Seek(0, 0); err != nil {
		t.Fatalf("failed to rewind temp file: %v", err)
	}
	os.Stdin = tempFile

	got := ReadInput()

	if got != want {
		t.Errorf("ReadInput() = %s, want %s", got, want)
	}
}

func TestGetVariant(t *testing.T) {
	originalStdin := os.Stdin
	defer func() { os.Stdin = originalStdin }()

	tempFile, err := os.CreateTemp("", "mock_stdin")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer func(name string) {
		_ = os.Remove(name)
	}(tempFile.Name())

	want := "3"
	if _, err := tempFile.Write([]byte(want + "\n")); err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}

	if _, err := tempFile.Seek(0, 0); err != nil {
		t.Fatalf("failed to rewind temp file: %v", err)
	}
	os.Stdin = tempFile

	input := ReadInput()

	variants := map[int]string{
		0: "variant1",
		1: "variant2",
		2: "variant3",
	}

	if _, err := GetVariant(input, variants); err != nil {
		t.Errorf("GetVariant() err != nil, expected non-error")
	}
}

func TestGetVariant_NonNumeric(t *testing.T) {
	originalStdin := os.Stdin
	defer func() { os.Stdin = originalStdin }()

	tempFile, err := os.CreateTemp("", "mock_stdin")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer func(name string) {
		_ = os.Remove(name)
	}(tempFile.Name())

	want := "test"
	if _, err := tempFile.Write([]byte(want + "\n")); err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}

	if _, err := tempFile.Seek(0, 0); err != nil {
		t.Fatalf("failed to rewind temp file: %v", err)
	}
	os.Stdin = tempFile

	input := ReadInput()

	variants := map[int]string{
		0: "variant1",
		1: "variant2",
		2: "variant3",
	}

	if _, err := GetVariant(input, variants); err == nil {
		t.Errorf("GetVariant() err = nil, want error")
	}
}
