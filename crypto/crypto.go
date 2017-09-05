package crypto

import (
	"crypto/rsa"
	"crypto/rand"
	"crypto"
	"fmt"
	"crypto/sha256"
	"encoding/base64"
)

// This file implements common cryptographic functions used to sign and verify the content(e.g. kubernetes manifests).

// Sign calculated the RSA signature of hashed value using RSA private key.
//
// To generate hash you can use generate.SHA256Hash function.
func Sign(privateKey *rsa.PrivateKey, hash crypto.Hash, hashed []byte) ([]byte, error){
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, hash, hashed)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return signature, nil
}

// VerifySignature verifies the RSA signature using RSA public key and hashed value.
// Note that you should calculate hash again in order to prevent any modification in transport (MITM attack).
//
// To generate hash you can use generate.SHA256Hash function.
func VerifySignature(publicKey *rsa.PublicKey, hash crypto.Hash, hashed []byte, signature []byte) bool {
	err := rsa.VerifyPKCS1v15(publicKey, hash, hashed, signature)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// SHA256Hash computes SHA256 hash of data and returns hash in []byte format.
func SHA256Hash(data []byte) []byte {
	h := sha256.New()
	h.Write([]byte(data))
	hashed := h.Sum(nil)
	return hashed
}

// SHA256ToBase64 encoded hashed value to base64 format.
func SHA256ToBase64(hashed []byte) string {
	return base64.URLEncoding.EncodeToString(hashed)
}