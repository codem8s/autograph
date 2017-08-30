package crypto

import (
	"crypto/rsa"
	"crypto/rand"
	"crypto"
	"fmt"
)

func Sign(privateKey *rsa.PrivateKey, hash crypto.Hash, hashed []byte) ([]byte, error){
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, hash, hashed)
	if err != nil {
		fmt.Println(err.Error)
		return nil, err
	}
	return signature, nil
}

func VerifySignature(publicKey *rsa.PublicKey, hash crypto.Hash, hashed []byte, signature []byte) bool {
	err := rsa.VerifyPKCS1v15(publicKey, hash, hashed, signature)
	if err != nil {
		fmt.Println(err.Error)
		return false
	}
	return true
}
