/*
Copyright 2017 Codem8s.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package crypto

import (
	"crypto"
	"testing"

	"github.com/codem8s/autograph/generate"
)

func TestVerifySignature(t *testing.T) {
	// generate rsa keypair
	privateKey, publicKey, err := generate.KeyPair()
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
