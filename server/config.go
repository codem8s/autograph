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

package server

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
)

func configTLS(certificatesDirectory string) (*tls.Config, error) {
	caCert, err := ioutil.ReadFile(certificatesDirectory + "/ca.pem")
	if err != nil {
		log.Fatal(err)
	}

	apiserverCA := x509.NewCertPool()
	apiserverCA.AppendCertsFromPEM(caCert)

	sCert, err := tls.LoadX509KeyPair(certificatesDirectory+"/server.pem", certificatesDirectory+"/server.key")
	if err != nil {
		return nil, err
	}
	return &tls.Config{
		Certificates: []tls.Certificate{sCert},
		ClientCAs:    apiserverCA,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}, nil
}
