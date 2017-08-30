package crypto

import (
	"testing"
	"github.com/codem8s/autograph/generate"
	"crypto/sha256"
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

	// generate signature from hash
	signature, err := Sign(privateKey, crypto.SHA256, hash("Super secret message"))
	if err != nil {
		t.Error("Cannot generate SHA256 hash")
	}

	// verify signature
	verified := VerifySignature(publicKey, crypto.SHA256, hash("Super secret message"), signature)
	if !verified {
		t.Error("Signature verification failed")
	}
}

func hash(data string) []byte {
	hash := sha256.New()
	hash.Write([]byte(data))
	hashed := hash.Sum(nil)
	return hashed
}
