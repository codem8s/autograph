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
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

// This file implements necessary functions to generate and save RSA key pair.

// PrivateKeyFile - the private key filename
// PublicKeyFile  - the public key filename
const (
	PrivateKeyFile = "autograph.key"
	PublicKeyFile  = "autograph.pub"
)

func generatePrivateKey(bits int) (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return privateKey, nil
}

func generatePublicKey(privateKey *rsa.PrivateKey) *rsa.PublicKey {
	publicKey := &privateKey.PublicKey
	return publicKey
}

func savePrivateKey(privateKey *rsa.PrivateKey) error {
	PrivASN1 := x509.MarshalPKCS1PrivateKey(privateKey)

	privBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: PrivASN1,
	})
	ioutil.WriteFile(PrivateKeyFile, privBytes, 0600)

	return nil
}

func savePublicKey(publicKey *rsa.PublicKey) error {
	PubASN1, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		fmt.Println(err)
		return err
	}

	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: PubASN1,
	})
	ioutil.WriteFile(PublicKeyFile, pubBytes, 0600)

	return nil
}

// KeyPair generated RSA key pair with 2048 bits length.
// Also writes RSA key pair to the file system
// under names defined in PrivateKeyFile and PublicKeyFile constants.
//
// Function takes *rsa.PrivateKey and *rsa.PublicKey as arguments
// It might be parsed/loaded using following functions:
// - keys.LoadPrivateKey
// - keys.LoadPublicKey
func KeyPair() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	fmt.Println("Generating private and public RSA key pair..")

	privateKey, err := generatePrivateKey(2048)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	publicKey := generatePublicKey(privateKey)

	err = savePrivateKey(privateKey)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	err = savePublicKey(publicKey)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	fmt.Println("Successfully generated private and public RSA key pair")
	return privateKey, publicKey, nil
}
