# Autograph (pre-alpha)
[![Build Status](https://travis-ci.org/codem8s/autograph.svg?branch=master)](https://travis-ci.org/codem8s/autograph)
![Version](https://img.shields.io/badge/version-0.0.1-brightgreen.svg)
[![Docker Repository on Quay.io](https://quay.io/repository/codem8s/autograph/status "Docker Repository on Quay.io")](https://quay.io/repository/codem8s/autograph)
[![Coverage](https://codecov.io/gh/codem8s/autograph/branch/master/graph/badge.svg "Test Coverage")](https://codecov.io/gh/codem8s/autograph)
[![Go Report Card](https://goreportcard.com/badge/github.com/codem8s/autograph "Go Report Card")](https://goreportcard.com/report/github.com/codem8s/autograph)
[![GoDoc](https://godoc.org/github.com/codem8s/autograph?status.svg "GoDoc Documentation")](https://godoc.org/github.com/codem8s/autograph)

Certificate signer and custom admission controller for Kubernetes manifests.

## Usage

    NAME:
      autograph - A new cli application

    USAGE:
      autograph [global options] command [command options] [arguments...]

    VERSION:
      0.0.1

    COMMANDS:
      generate, g  generate a key and certificate pair
      sign, s      sign a manifest
      verify, v    verify a signed manifest
      run, v       starts the HTTP(S) server
      help, h      Shows a list of commands or help for one command

    GLOBAL OPTIONS:
      --help, -h     show help
      --version, -v  print the version
      
 To sign your manifest:
 
      autograph sign example-manifest.yaml
    
After that there should be a new annotation in the manifest, e.g.:

    ...
    metadata:
      annotations:
        autograph.codemat.es/signature: 72976B7400E7630F846501847CB04A...
    ...

### Commands:
- generate - generate a key and certificate pair
- sign - sign a manifest
- verify - verify a signed manifest
- run - starts the HTTP(S) server

### Dependencies

- Go 1.8.0+
- Kubernetes 1.7.0+

## Flow

1. Signer (CLI tool) signs a manifest using a provided key an puts the signature in the manifest.
2. Verifier (an admission controller) checks the signature with a provided certificate.
3. If the signature is correct the manifest is deployed (or more precisely, it's is handed over to other admission controllers).

## Build from source code

### Define go workspace (GOPATH)

    export GOPATH=~/go
    
### Get the repository
    
    go get -u github.com/codem8s/autograph
    cd $GOPATH/src/github.com/codem8s/autograph    

### Build

    go build
    
### Run tests

    go test

## Run on minikube

### Installation for Ubuntu

    sudo apt-get update
    sudo apt-get install virtualbox
    curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
    chmod +x ./kubectl
    sudo mv ./kubectl /usr/local/bin/kubectl
    curl -Lo minikube https://storage.googleapis.com/minikube/releases/v0.22.0/minikube-linux-amd64 && chmod +x minikube && sudo mv minikube /usr/local/bin/

### Run

    cd ~/go
    export GOPATH=$(pwd)
    cd $GOPATH/src/github.com/codem8s/autograph
    ./gencerts.sh
    ./start-minikube.sh
    export CGO_ENABLED=0 GOOS=linux
    go build
    eval $(minikube docker-env)
    docker build -t autograph .
    kubectl create -f kubernetes/service.yaml
    kubectl create -f kubernetes/autograph.yaml
    
### Test

    kubectl create -f kubernetes/echoserver.yaml
    kubectl get po
    kubectl logs autograph

## Dependency management
    
### Installation    
    
    cd ~/go
    go get -u github.com/golang/dep/cmd/dep
    
### Usage

    export GOPATH=$(pwd)
    export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
    cd $GOPATH/src/github.com/codem8s/autograph
    dep ensure

## Contribute

If you have any idea for an improvement or found a bug don't hesitate to open an issue or just make a pull request!

## Useful links

- [Kubernetes Admission Controllers](https://kubernetes.io/docs/admin/extensible-admission-controllers)
- [Kubernetes Initializer Tutorial](https://github.com/kelseyhightower/kubernetes-initializer-tutorial)
- [Example admission controller](https://github.com/caesarxuchao/example-webhook-admission-controller)
- [Golang on Ubuntu](https://github.com/golang/go/wiki/Ubuntu)
- [Golang Dep](https://github.com/golang/dep)
