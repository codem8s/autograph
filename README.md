# Autograph (pre-alpha)
Certificate signer and custom admission controller for Kubernetes manifests.

## Usage
To sign your manifest:

  autograph sign example-manifest.yaml
  
After that there should be a new annotation in the manifest, e.g.:

    ...
      annotations:
        autograph.codemat.es/signature: 72976B7400E7630F846501847CB04A...
    ...

## Flow

1. Signer (CLI tool) signs a manifest using a provided key an puts the signature in the manifest.
2. Verifier (an admission controller) checks the signature with a provided certificate.
3. If the signature is correct the manifest is deployed (or more precisly, it's is handed over to other admission controllers).

## Useful links

[1] https://kubernetes.io/docs/admin/extensible-admission-controllers/

[2] https://github.com/kelseyhightower/kubernetes-initializer-tutorial

[3] https://github.com/caesarxuchao/example-webhook-admission-controller
