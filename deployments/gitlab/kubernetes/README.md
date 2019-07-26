# GitLab on Kubernetes

This is a self-contained project that installs GitLab onto a Kubernetes cluster using Helm.

# Pre-Requisites

1. Kubectl
2. Minikube
3. Helm

Installation recipes can be found in the [`~/init` directory](../../../init).

# Get Started

## Kubernetes

Start a minikube cluster by running:

```sh
make minikube
```

## Kubernetes Namespace

To create the required namespaces, services accounts, and roles, run:

```sh
make kubernetes
```

## Helm/Tiller

Navigate into the `./helm` directory and continue from the [`README.md` there](./helm/README.md)

