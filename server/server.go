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
	"encoding/json"
	"io/ioutil"
	"net/http"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/pkg/api/v1"

	"strings"
	"log"
)

// only allow pods to pull images from specific registry.
func admit(data []byte) *AdmissionReviewStatus {
	ar := AdmissionReview{}
	if err := json.Unmarshal(data, &ar); err != nil {
		log.Print(err)
		return nil
	}
	// The externalAdmissionHookConfiguration registered via selfRegistration
	// asks the kube-apiserver only sends admission request regarding pods.
	podResource := metav1.GroupVersionResource{Group: "", Version: "v1", Resource: "pods"}
	if ar.Spec.Resource != podResource {
		log.Printf("expect resource to be %s\n", podResource)
		return nil
	}

	raw := ar.Spec.Object.Raw
	pod := v1.Pod{}
	if err := json.Unmarshal(raw, &pod); err != nil {
		log.Print(err)
		return nil
	}
	reviewStatus := AdmissionReviewStatus{}
	for _, container := range pod.Spec.Containers {
		// gcr.io is just an example.
		if !strings.Contains(container.Image, "gcr.io") {
			reviewStatus.Allowed = false
			reviewStatus.Result = &metav1.Status{
				Reason: "can only pull image from grc.io",
			}
			return &reviewStatus
		}
	}
	reviewStatus.Allowed = true
	return &reviewStatus
}

func serve(w http.ResponseWriter, r *http.Request) {
	var body []byte
	if r.Body != nil {
		if data, err := ioutil.ReadAll(r.Body); err == nil {
			body = data
		}
	}
	log.Printf("Review request:\n%s", body)

	// verify the content type is accurate
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		log.Printf("contentType=%s, expect application/json", contentType)
		return
	}

	reviewStatus := admit(body)
	ar := AdmissionReview{
		Status: *reviewStatus,
	}

	resp, err := json.Marshal(ar)
	if err != nil {
		log.Print(err)
	}
	if _, err := w.Write(resp); err != nil {
		log.Print(err)
	}
}

// Run starts http server which verifies incoming manifests in kubernetes
func Run(certificatesDirectory string) error {
	log.Print("Server is starting ...")
	http.HandleFunc("/", serve)
	tlsConfig, err := configTLS(certificatesDirectory)
	if err != nil {
		return err
	}
	server := &http.Server{
		Addr:      ":8000",
		TLSConfig: tlsConfig,
	}
	return server.ListenAndServeTLS("", "")
}
