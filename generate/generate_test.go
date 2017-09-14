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

package generate

import (
	"os"
	"testing"
)

func TestGenerateRSAKeyPair(t *testing.T) {
	// when
	privateKey, publicKey, err := KeyPair()

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
