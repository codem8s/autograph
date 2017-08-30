package generate

import (
	"testing"
)

func TestVerifySignature(t *testing.T) {
	_, err := LoadPrivateKey("../resources/autograph.key")
	if err != nil {
		t.Error("Cannot load private key")
	}

	_, err = LoadPublicKey("../resources/autograph.pub")
	if err != nil {
		t.Error("Cannot load public key")
	}
}