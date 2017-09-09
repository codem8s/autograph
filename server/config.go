package server

import (
	"crypto/tls"
	"crypto/x509"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/apis/admissionregistration/v1alpha1"
	"k8s.io/client-go/rest"

	"os"

	"github.com/golang/glog"
)

const (
	projectName = "autograph"
)

// get a clientset with in-cluster config.
func getClient() *kubernetes.Clientset {
	config, err := rest.InClusterConfig()
	if err != nil {
		glog.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		glog.Fatal(err)
	}
	return clientset
}

func configTLS(caCert []byte) *tls.Config {
	//cert := getAPIServerCert(clientset)
	apiserverCA := x509.NewCertPool()
	apiserverCA.AppendCertsFromPEM(caCert)

	sCert, err := tls.X509KeyPair(serverCert, serverKey)
	if err != nil {
		glog.Fatal(err)
	}
	return &tls.Config{
		Certificates: []tls.Certificate{sCert},
		ClientCAs:    apiserverCA,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}
}

// register this example webhook admission controller with the kube-apiserver
// by creating externalAdmissionHookConfigurations.
func selfRegistration(clientset *kubernetes.Clientset, caCert []byte) {
	//time.Sleep(10 * time.Second)
	client := clientset.AdmissionregistrationV1alpha1().ExternalAdmissionHookConfigurations()
	_, err := client.Get(projectName, metav1.GetOptions{})
	if err == nil {
		if err2 := client.Delete(projectName, nil); err2 != nil {
			glog.Fatal(err2)
		}
	}
	webhookConfig := &v1alpha1.ExternalAdmissionHookConfiguration{
		ObjectMeta: metav1.ObjectMeta{
			Name: projectName,
		},
		ExternalAdmissionHooks: []v1alpha1.ExternalAdmissionHook{
			{
				Name: "pod-image.k8s.io",
				Rules: []v1alpha1.RuleWithOperations{{
					Operations: []v1alpha1.OperationType{v1alpha1.Create, v1alpha1.Update},
					Rule: v1alpha1.Rule{
						APIGroups:   []string{""},
						APIVersions: []string{"v1"},
						Resources:   []string{"pods"},
					},
				}},
				ClientConfig: v1alpha1.AdmissionHookClientConfig{
					Service: v1alpha1.ServiceReference{
						Namespace: os.Getenv("NAMESPACE"),
						Name:      projectName,
					},
					CABundle: caCert,
				},
			},
		},
	}
	if _, err := client.Create(webhookConfig); err != nil {
		glog.Fatal(err)
	}
}
