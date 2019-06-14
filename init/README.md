# `init` (WIP)

This directory contains setup scripts for common cloud-native support services. Services here have been set up for deploying on a single-node Kubernetes (Minikube) for the purposes of development/testing before deploying to production. Production tuning is left as an exercise for the reader.

# Assumptions

These scripts were created for deployment in both a local development and production environment with the following assumptions:

**Technical**
1. Kubernetes API server is exposed to the machine running these
1. Helm (with a TLS/namespace-secured Tiller) is used to deploy all services

# Pre-Requisites

- [Docker]()
- [Docker Compose]()
- Hypervisor
  - MacOS: [HyperKit](https://github.com/kubernetes/minikube/blob/master/docs/drivers.md#hyperkit-driver)
  - Linux: [KVM2](https://github.com/kubernetes/minikube/blob/master/docs/drivers.md#kvm2-driver)
  - All: [VirtualBox]()

# Order of Installation

## Group 1 (Platform-level)

- [Minikube](./minikube)

## Group 2 (Release-level)

- [Helm](./helm)

## Group 3 (Infra-level)

- [Istio](./istio)
- [LinkerD](./linkerd)
- [Prometheus Operator](./prometheus-operator)
- [Vault with Key-Value storage](./vault-key-value)
- [Vault with Dynamic Secrets storage](./vault-dynamic-secrets)

## Group 4 (App-level)

- [Knative](./knative)
- [Showcase Deployments](../deployments/showcase/kubernetes)

# Types of Services

## Development

### Services

- [Minikube](#minikube)

## Deployment Manager

### Services

- [Helm](#helm)

## Secrets 

### Services

- [Vault](#vault)

## Service Mesh

### Services

- [Istio](#istio)
- [LinkerD](#linkerd)

## System Observability

### Services

- [Prometheus](#prometheus)
- [Jaegar](#jaegar)
- [Zipkin](#zipkin)

## Telemetry Visualisation

Telemetry visuali

### Services Reference

- [Kiali](#kiali)
- [Grafana](#grafana)
- [Kibana](#kibana)

# Services

## Helm

Helm is a service which facilitates deployment of your services. We use [Helm Charts](https://github.com/helm/charts) to deploy many of these services.

> [Check out Helm](https://helm.sh/)

## Istio

Istio is a service mesh which assists in the instrumentation of your cluster. The main benefits from Istio are:

1. Overview of service topology
2. Enabling of service telemetry
3. Enabling of network telemetry
  
> [Check out Istio](https://istio.io/)

## Knative

**(WIP)**

Knative is a Functions-as-a-Service (Faas) service that enables the serverless architecture pattern. KNative is useful for deployment of scheduled/one-off jobs that require access to services accessible from the network hosting the services.

> [Check out Knative](https://knative.dev/)

## LinkerD

LinkerD is a service mesh