package encryption

import (
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	password := "mysecretpassword"
	data := []byte("this is a test")

	encrypted, err := Encrypt(data, password)
	if err != nil {
		t.Fatalf("Encryption failed: %v", err)
	}

	decrypted, err := Decrypt(encrypted, password)
	if err != nil {
		t.Fatalf("Decryption failed: %v", err)
	}

	if string(decrypted) != string(data) {
		t.Errorf("Decrypted data does not match original data. Got %s, want %s", string(decrypted), string(data))
	}
}
