---
layout: page
title: KinD
permalink: installation/platforms/kind
---

# Quick Start with KinD

## Installation and usage

On Mac / Linux via Homebrew:
```
brew install kind
```

On macOS / Linux:
```
curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.8.1/kind-$(uname)-amd64
chmod +x ./kind
mv ./kind /some-dir-in-your-PATH/kind
```
If you are running Ubuntu on WSL2, use `Docker Ubuntu` distro to install `Docker`. 

### Create cluster using KinD

In order to successfully build the meshery server on your local machine, you will need to follow these specific instructions
according to your Operating System to complete the creation of KinD cluster.

#### KinD on WSL2

First, we will get the ip address of your WSL2 distro by:
```
ip addr | grep eth0
```

You will see the output like:
```
4: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc mq state UP group default qlen 1000
    inet 172.1.1.1/20 brd 172.1.1.255 scope global eth0
```

Copy the ip address, we will use that in the next step.


Then, create a file called `kind_cluster.yaml` and put the ip address under `apiServerAddress`:
```
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
networking:
  apiServerAddress: "172.1.1.1"
```

Now create the KinD cluster with the config file `kind_cluster.yaml`:

```
kind create cluster --config kind_cluster.yaml --name kind --wait 300s
```

You will see

```
Creating cluster "kind" ...
 • Ensuring node image (kindest/node:v1.17.0) 🖼  ...
 ✓ Ensuring node image (kindest/node:v1.17.0) 🖼
 • Preparing nodes 📦   ...
 ✓ Preparing nodes 📦 
 • Writing configuration 📜  ...
 ✓ Writing configuration 📜
 • Starting control-plane 🕹️  ...
 ✓ Starting control-plane 🕹️
 • Installing CNI 🔌  ...
 ✓ Installing CNI 🔌
 • Installing StorageClass 💾  ...
 ✓ Installing StorageClass 💾
 • Waiting ≤ 5m0s for control-plane = Ready ⏳  ...
 ✓ Waiting ≤ 5m0s for control-plane = Ready ⏳
 • Ready after 59s 💚
Set kubectl context to "kind-kind"
You can now use your cluster with:

kubectl cluster-info --context kind-kind

Not sure what to do next? 😅 Check out https://kind.sigs.k8s.io/docs/user/quick-start/
```

#### KinD on other systems

Creating a Kubernetes cluster is as simple as `kind create cluster`.

For more configuration of installtion, please refer to KinD official documentation.


### Access the KinD cluster

By default, the cluster access configuration is stored in ${HOME}/.kube/config if $KUBECONFIG environment variable is not set. You can set the `KIUBECONFIG` environment command below:
```
export KUBECONFIG=${HOME}/.kube/config
```
Use the command below check the connection of the cluster and make sure the cluster you connected what's the cluster was created by KinD:
```
kubectl cluster-info --context kind-kind
```

To delete your cluster use: 
```
kind delete cluster --name kind
```


<a name="helm"></a>
## Using Helm

### Helm v3

We strongly recommend to use Helm v3, because of the version not included the Tiller(https://helm.sh/blog/helm-3-preview-pt2/#helm) component anymore. It’s more light and safe.

Run the following:
```
$ git clone https://github.com/layer5io/meshery.git; cd meshery
$ kubectl create namespace meshery
$ helm install meshery --namespace meshery install/kubernetes/helm/meshery
```

* **NodePort** - If your cluster does not have an Ingress Controller or a load balancer, then use NodePort to expose Meshery and that can be modify under the chart `values.yaml`:

```
service:
  type: NodePort
  port: 8080
  annotations: {}
```
