package crypto

import (
	"crypto/rsa"
	"crypto/rand"
	"crypto"
	"fmt"
	"crypto/sha256"
	"encoding/base64"
)

func Sign(privateKey *rsa.PrivateKey, hash crypto.Hash, hashed []byte) ([]byte, error){
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, hash, hashed)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return signature, nil
}

func VerifySignature(publicKey *rsa.PublicKey, hash crypto.Hash, hashed []byte, signature []byte) bool {
	err := rsa.VerifyPKCS1v15(publicKey, hash, hashed, signature)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func SHA256Hash(data []byte) []byte {
	h := sha256.New()
	h.Write([]byte(data))
	hashed := h.Sum(nil)
	return hashed
}

func SHA256ToBase64(hashed []byte) string {
	return base64.URLEncoding.EncodeToString(hashed)
}