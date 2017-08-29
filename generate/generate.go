package generate

import (
	"fmt"
	"crypto/rsa"
	"crypto/rand"
	"crypto/x509"
	"io/ioutil"
	"encoding/pem"
)

func generatePrivateKey(bits int) (*rsa.PrivateKey, error){
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)

	if err != nil {
		fmt.Println(err.Error)
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
	ioutil.WriteFile("private.key", privBytes, 0644) // FIXME @antoniaklja shall we pass destination path here?

	return nil
}

func savePublicKey(publicKey *rsa.PublicKey) error{
	PubASN1, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		fmt.Println(err)
		return err
	}

	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: PubASN1,
	})
	ioutil.WriteFile("key.pub", pubBytes, 0644)  // FIXME @antoniaklja shall we pass destination path here?

	return nil
}

func Generate() error {
	fmt.Println("Generating private and public RSA key pair..")

	privateKey, err := generatePrivateKey(2048)
	if err != nil {
		fmt.Println(err.Error)
		return err
	}

	publicKey := generatePublicKey(privateKey)

	savePrivateKey(privateKey)
	savePublicKey(publicKey)

	fmt.Println("Successfully generated private and public RSA key pair")
	return nil
}