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
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"os"
	"time"
	"crypto/tls"
	"net"
	"log"
)

const (
	cnBase = "codem8s_autograph"
	certificateBits = 2048
	twoYears = 2
	CertificatesDestinationFolder = "resources/"
	caFile = CertificatesDestinationFolder + "ca.pem"
	caKeyFile = CertificatesDestinationFolder + "ca.key"
	fileMode = 0600
)

// GenerateTLSCertificates generate TLS certificates and put them in resources directory
func GenerateTLSCertificates() {
	os.MkdirAll(CertificatesDestinationFolder, 0700)
	generateCA()

	caPair, err := tls.LoadX509KeyPair(caFile, caKeyFile)
	if err != nil {
		log.Fatal(err)
	}
	// The subjectAltName/IP address in the certificate MUST match the one configured on the Kubernetes Service.
	kubernetesServiceIP := []net.IP{net.ParseIP("10.0.0.231")}

	generateCertificate("client", caPair, kubernetesServiceIP)
	generateCertificate("server", caPair, kubernetesServiceIP)
}

func generateSerialNumber() (*big.Int) {
	serialNumber, err := rand.Int(rand.Reader, (&big.Int{}).Exp(big.NewInt(2), big.NewInt(159), nil))
	if err != nil {
		log.Fatal(err)
	}
	return serialNumber
}

func generateCA() {
	serialNumber := generateSerialNumber()

	ca := &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName:  cnBase + "_ca",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(twoYears, 0, 0),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, certificateBits)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := &privateKey.PublicKey
	caBinary, err := x509.CreateCertificate(rand.Reader, ca, ca, publicKey, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	certOut, err := os.Create(caFile)
	if err != nil {
		log.Fatal(err)
	}
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: caBinary})
	err = certOut.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Print(caFile + " created \n")

	// Private key
	keyOut, err := os.OpenFile(caKeyFile, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, fileMode)
	if err != nil {
		log.Fatal(err)
	}
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)})
	err = keyOut.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Print(caKeyFile + " created \n")
}

func generateCertificate(name string, caPair tls.Certificate, kubernetesServiceIP []net.IP) {
	ca, err := x509.ParseCertificate(caPair.Certificate[0])
	if err != nil {
		log.Fatal(err)
	}

	serialNumber := generateSerialNumber()

	// Prepare certificate
	cert := &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName:  cnBase + "_" + name,
		},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(twoYears, 0, 0),
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
		IPAddresses: kubernetesServiceIP,
		IsCA: false,
	}
	priv, err := rsa.GenerateKey(rand.Reader, certificateBits)
	if err != nil {
		log.Fatal(err)
	}
	pub := &priv.PublicKey

	// Sign the certificate
	certBinary, err := x509.CreateCertificate(rand.Reader, cert, ca, pub, caPair.PrivateKey)

	// Public key
	pemFile := CertificatesDestinationFolder + name + ".pem"
	certOut, err := os.Create(pemFile)
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: certBinary})
	certOut.Close()
	log.Print(pemFile + " created \n")

	// Private key
	keyFile := CertificatesDestinationFolder + name + ".key"
	keyOut, err := os.OpenFile(keyFile, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, fileMode)
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})
	keyOut.Close()
	log.Print(keyFile + " created \n")
}