package generate

import (
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"encoding/pem"
)

func LoadPrivateKey(path string) (*rsa.PrivateKey, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	block, _ := pem.Decode([]byte(bytes))
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return privateKey, nil
}

func LoadPublicKey(path string) (*rsa.PublicKey, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	block, _ := pem.Decode([]byte(bytes))
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	rsaPubKey := publicKey.(*rsa.PublicKey)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return rsaPubKey, nil
}