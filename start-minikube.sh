#!/usr/bin/env bash

# Mounted Host Folder
# Driver	    OS	    HostFolder	VM
#VirtualBox	    Linux	/home	    /hosthome
# Set to directory containing generated TLS certificates.
CERT_DIR=`echo $(pwd) | sed -e 's/home/hosthome/g'`
echo "Certs directory '${CERT_DIR}'"

# Set to admission controllers to include. This example uses the default set of
# admission controllers enabled in the Kubernetes API server plus the
# GenericAdmissionWebhook admission controller.
ADMISSION_CONTROLLERS=NamespaceLifecycle,LimitRanger,ServiceAccount,PersistentVolumeLabel,DefaultStorageClass,GenericAdmissionWebhook,ResourceQuota,DefaultTolerationSeconds

minikube start --kubernetes-version v1.7.0 \
    --extra-config=apiserver.Admission.PluginNames=$ADMISSION_CONTROLLERS \
    --extra-config=apiserver.ProxyClientCertFile=$CERT_DIR/clientCert.pem \
    --extra-config=apiserver.ProxyClientKeyFile=$CERT_DIR/clientKey.pem