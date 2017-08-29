package generate

import (
	"testing"
	"os"
)

func TestGenerateRSAKeyPair(t *testing.T) {
	// when
	GenerateKeyPair()

	// then
	if _, err := os.Stat(PrivateKeyFile); os.IsNotExist(err) {
		t.Error("Private Key doesn't exist")
	} else {
		os.Remove(PrivateKeyFile)
	}

	if _, err := os.Stat(PublicKeyFile); os.IsNotExist(err) {
		t.Error("Public Key doesn't exist")
	} else {
		os.Remove(PublicKeyFile)
	}
	// FIXME @antoniaklja add more versatile tests
}
