package generate

import (
	"testing"
	"os"
)

func TestGenerateRSAKeyPair(t *testing.T) {
	// when
	privateKey, publicKey, err := GenerateKeyPair()

	// then
	if err != nil {
		t.Error("Cannot generate RSA keypair")
	}

	if privateKey == nil {
		t.Error("Private Key doesn't exist")
	}

	if publicKey == nil {
		t.Error("Public Key doesn't exist")
	}

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

	// TODO @antoniaklja add private and public keys integrity check
}