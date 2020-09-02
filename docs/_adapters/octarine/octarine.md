---
layout: page
title: Octarine Adapter
name: Octarine
version: v1.0
port: 10003/tcp
project_status: stable
github_link: https://github.com/layer5io/meshery-octarine
image: /docs/assets/img/service-meshes/octarine.svg
---
# {{ page.name }}

| Service Mesh   | Adapter Status | Latest Supported Mesh Version |
| :------------: | :------------:   | :------------:              |
| {{page.title}} | [{{ page.project_status }}]({{ page.github_link }}) | {{page.version}}  |

### Lifecycle management

The {{page.name}} can install **{{page.version}}** of the {{page.name}} service mesh. The SMI adapter for Kuma can also be installed using Meshery.

### Features

1. Policy-based validation that k8s workloads spec are secure.
1. Visibility of layer 4-7 traffic between workloads, as well as ingress and egress.
1. Encryption and authentication based on mTLS.
1. Automation and enforcement of access control policy based on observed traffic.
1. Threat detection based on signatures and anomalies.

### Configuration
In order to connect to the Octarine Control Plane, the adapter requires the following environment variables to be set:

* **OCTARINE_DOCKER_USERNAME** : The docker username needed to pull Octarine's images to the target cluster, supplied by Octarine. Do not use your own docker credentials.
* **OCTARINE_DOCKER_EMAIL** : The docker email, supplied by Octarine.
* **OCTARINE_DOCKER_PASSWORD** : The docker password, supplied by Octarine.
* **OCTARINE_ACC_MGR_PASSWD** : The password that will be assigned to the user 'meshery' in the new account.
* **OCTARINE_CREATOR_PASSWD** : The password needed to create an account in Octarine.
* **OCTARINE_DELETER_PASSWD** : The password needed to delete the account in Octarine.
* **OCTARINE_CP** : The address of the Octarine Control Plane. `Example: meshery-cp.octarinesec.com`
* **OCTARINE_DOMAIN** : The name that will be assigned to the target cluster in Octarine. `Example: meshery:domain`


### Usage

Once the Octarine's data plane services are deployed, the adapter can be used to deploy Bookinfo:

* Enable the target namespace for automatic sidecar injection.
* Deploy Bookinfo to the target namespace.

### Architecture

#### Control Plane

![Alt text](./octarine_cparch.jpg?raw=true "Octarine Control Plane")

#### Data Plane

![Alt text](./octarine_dparch.jpg?raw=true "Octarine Data Plane")

### Suggested Topics

- Examine [Meshery's architecture]({{ site.baseurl }}/architecture) and how adapters fit in as a component.
- Learn more about [Meshery Adapters]({{ site.baseurl }}/architecture/adapters).

### Sample Applications 

The Meshery adapter for {{ page.name }} includes the below sample application operation. Meshery can be use to deploy this sample application.

- [Istio BookInfo](https://github.com/layer5io/istio-service-mesh-workshop/blob/master/lab-2/README.md#what-is-the-bookinfo-application)
    - This application is a polyglot composition of microservices are written in different languages and sample BookInfo application displays information about a book, similar to a single catalog entry of an online book store.
