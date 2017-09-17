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
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestGenerateTLSCertificates(t *testing.T) {
	// when
	TLSCertificates()

	// then
	fileName := CertificatesDestinationFolder + "ca.pem"
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		t.Error("CA certificate doesn't exist")
	}

	fileName = CertificatesDestinationFolder + "ca.key"
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		t.Error("CA key doesn't exist")
	}

	fileName = CertificatesDestinationFolder + "client.pem"
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		t.Error("Client certificate doesn't exist")
	}

	fileName = CertificatesDestinationFolder + "client.key"
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		t.Error("Client key doesn't exist")
	}

	fileName = CertificatesDestinationFolder + "server.pem"
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		t.Error("Server certificate doesn't exist")
	}

	fileName = CertificatesDestinationFolder + "server.key"
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		t.Error("Server key doesn't exist")
	}
}

func TestLoadTLSCertificates(t *testing.T) {
	// when
	TLSCertificates()

	// then
	_, err := tls.LoadX509KeyPair(CertificatesDestinationFolder+"ca.pem", CertificatesDestinationFolder+"ca.key")
	if err != nil {
		t.Error("CA certificate can't be load")
	}

	_, err = tls.LoadX509KeyPair(CertificatesDestinationFolder+"client.pem", CertificatesDestinationFolder+"client.key")
	if err != nil {
		t.Error("Client certificate can't be load")
	}

	_, err = tls.LoadX509KeyPair(CertificatesDestinationFolder+"server.pem", CertificatesDestinationFolder+"server.key")
	if err != nil {
		t.Error("Server certificate can't be load")
	}
}

func TestVerifyTLSCertificates(t *testing.T) {
	// when
	TLSCertificates()

	// then
	verifyCertificate(CertificatesDestinationFolder+"ca.pem", CertificatesDestinationFolder+"client.pem")
	verifyCertificate(CertificatesDestinationFolder+"ca.pem", CertificatesDestinationFolder+"server.pem")
}

func verifyCertificate(caFile, certificateFile string) {
	caCertificate, err := ioutil.ReadFile(caFile)
	if err != nil {
		panic(fmt.Sprintf("Failed to load %v: %v", caFile, err))
	}
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(caCertificate)
	if !ok {
		panic(fmt.Sprintf("Failed to parse %v: %v", caFile, err))
	}

	certificate, err := ioutil.ReadFile(certificateFile)
	if err != nil {
		panic(fmt.Sprintf("Failed to load %v: %v", certificateFile, err))
	}
	block, _ := pem.Decode(certificate)
	if block == nil {
		panic(fmt.Sprintf("Failed to parse %v", certificateFile))
	}
	cert, errParse := x509.ParseCertificate(block.Bytes)
	if errParse != nil {
		panic(fmt.Sprintf("Failed to parse %v: %v", certificateFile, errParse))
	}

	opts := x509.VerifyOptions{
		Roots: roots,
	}

	if _, err := cert.Verify(opts); err != nil {
		panic(fmt.Sprintf("Failed to verify %v: %v", certificateFile, err))
	}
}
