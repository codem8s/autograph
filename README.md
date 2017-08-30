# Autograph (pre-alpha)
Certificate signer and custom admission controller for Kubernetes manifests.

## Usage

    NAME:
      autograph - A new cli application

    USAGE:
      autograph [global options] command [command options] [arguments...]

    VERSION:
      0.1

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

## Flow

1. Signer (CLI tool) signs a manifest using a provided key an puts the signature in the manifest.
2. Verifier (an admission controller) checks the signature with a provided certificate.
3. If the signature is correct the manifest is deployed (or more precisly, it's is handed over to other admission controllers).

## Useful links

- [Kubernetes Admission Controllers](https://kubernetes.io/docs/admin/extensible-admission-controllers)
- [Kubernetes Initializer Tutorial](https://github.com/kelseyhightower/kubernetes-initializer-tutorial)
- [Example admission controller](https://github.com/caesarxuchao/example-webhook-admission-controller)
