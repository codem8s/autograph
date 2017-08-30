package crypto

import (
	"testing"
	"github.com/codem8s/autograph/generate"
	"crypto"
)

func TestVerifySignature(t *testing.T) {
	// generate rsa keypair
	privateKey, publicKey, err := generate.GenerateKeyPair()
	if err != nil {
		t.Error("Cannot generate RSA keypair")
	}

	if privateKey == nil {
		t.Error("Private Key doesn't exist")
	}

	if publicKey == nil {
		t.Error("Public Key doesn't exist")
	}

	// generate hash and corresponding signature
	hashed := SHA256Hash([]byte("Super secret message"))
	signature, err := Sign(privateKey, crypto.SHA256, hashed)
	if err != nil {
		t.Error("Cannot generate SHA256 hash")
	}

	// generate hash again and verify signature
	hashedAgain := SHA256Hash([]byte("Super secret message"))
	verified := VerifySignature(publicKey, crypto.SHA256, hashedAgain, signature)
	if !verified {
		t.Error("Signature verification failed")
	}
}

func TestSHA256Hash(t *testing.T) {
	message := "super secret message"
	messageBase64SHA256Hash := "6f9oReRJXFLDPrUnozlPZ8sGI82nmBbct8R3phBBuaw="

	hashed := SHA256Hash([]byte(message))
	hashedBase64 := SHA256ToBase64(hashed)

	if messageBase64SHA256Hash != hashedBase64 {
		t.Error("Invalid SHA256 hash")
	}
}