---
layout: page
title: EKS
permalink: installation/platforms/eks
---

# Quick Start with Amazon Elastic Kubernetes Service (EKS)

In order to provide Meshery with the necessary access to your managed Kubernetes instance, 
Meshery will need to be assigned a `ServiceAccount`. An existing `ServiceAccount` can be used or a new one created. Ensure that the `ServiceAccount` you use has the `cluster-admin` role assigned.

_Note: Make sure you are able to access EKS with `kubectl` by following the [EKS Guide.](https://docs.aws.amazon.com/eks/latest/userguide/create-kubeconfig.html){:target="_blank"}_

## Automatic Configuration (Recommended)

1. In your browser, navigate to Meshery (e.g., `http://localhost:9081`) and login.
1. Download your Meshery authentication token by clicking **Get Token** under your user profile.
1. Use this authentication token to execute the following command:
    ```
    $ mesheryctl system config eks --token <PATH TO TOKEN>
    ```

This command updates your kubeconfig to provide Meshery with access to your managed Kubernetes instance.
Once configured, proceed with using Meshery (`mesheryctl system start`).

## Manual Configuration (Optional)

If the [Automatic Configuration](#automatic-configuration-recommended) procedure fails or you would like to manually prepare your kubeconfig file to provide Meshery with the necessary access to your managed Kubernetes instance, perform the following actions:

1. Create a `ServiceAccount` with `cluster-admin` role
    ```    
    $ kubectl create serviceaccount meshery
    ```
1. Adding/Binding `cluster-admin` role to new service account `meshery`
    ```
    $ kubectl create clusterrolebinding meshery-binding --clusterrole=cluster-admin \
    --serviceaccount=default:meshery
    ```
1. Get secret name from `ServiceAccount`.
    ```
    $ kubectl get secrets

    NAME                           TYPE                                  DATA   AGE
    default-token-fnfjp            kubernetes.io/service-account-token   3      95d
    meshery-token-5z9xj               kubernetes.io/service-account-token   3      66m
    ```
    _Note: Here the secret name is **meshery-token-5z9xj**_
1. Get secret/token
    ```
    $ kubectl describe secret  sa-1-token-5z9xj
    Name:         meshery-token-5z9xj
    Namespace:    default
    Labels:       <none>
    Annotations:  kubernetes.io/service-account.name: meshery
                  kubernetes.io/service-account.uid: 397XXX-XXX-XXXX-XXXXX-XXXXX

    Type:  kubernetes.io/service-account-token

    Data
    ====
    ca.crt:     1025 bytes
    namespace:  7 bytes
    token:      XXXhbGciOiJSUXXXX
    ```

1. Generate new kubeconfig yaml file to use as input to Meshery.
1. Set config Credential using above generate `token`.
    ```
    $ kubectl config set-credentials meshery --token=XXXXX

    o/p:User "meshery" set.
    ```
1. Set current context to our new service account `meshery`
    ```
    $ kubectl config set-context --current --user=meshery

    o/p:
    Context "aws" modified.
    ```
1. Generate kubeconfig yaml file to use as input to Meshery.
     ```
    $ kubectl config view --minify --flatten >  config_aws_eks.yaml
    ```
