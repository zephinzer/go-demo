define HELP_OUTPUT

Vault on Kubernetes
^^^^^^^^^^^^^^^^^^^

Vault is a secrets storage management tool. This deployment
of Vault demonstrates putting Vault on a Kubernetes cluster
using Consul as a backend. Use an external load balancer
to access the Vault service so that Kubernetes becomes just
an implementation detail.

Ref01: https://testdriven.io/blog/running-vault-and-consul-on-kubernetes

endef

include ../../../scripts/Makefile

prereq:
	#
	# get the go dependencies
	@ls -al $$(which go)
	#
	#
	# download the certificate creator tools - better than generating them by hand via opeenssl gen
	#
	@go get -v -u github.com/cloudflare/cfssl/cmd/cfssl
	@go get -v -u github.com/cloudflare/cfssl/cmd/cfssljson
	#
	# check that go get installed cfssl in the right place
	#
	@which cfssl && ls -al $$(which cfssl)
	@which cfssljson && ls -al $$(which cfssljson)
	#
	# check for minikube availability
	#
	@which minikube && ls -al $$(which minikube)
	#
	# consul is used for generation of the gossip encryption key
	#
	@which consul && ls -al $$(which consul)

# you need to run this before any of the `kubectl` commands will work
start_minikube:
	minikube start -p vault --vm-driver kvm2 --disk-space 5G --memory 2048 --cpus 2

######################
# secrets generation #
######################

secrets_create_ca:
	cfssl gencert -initca ./resources/cfssl/ca-csr.json | cfssljson -bare ./resources/certs/ca
	cfssl gencert \
		-ca=./resources/certs/ca.pem \
		-ca-key=./resources/certs/ca-key.pem \
		-config=./resources/cfssl/ca.json \
		-profile=default \
		./resources/cfssl/consul-csr.json | cfssljson -bare ./resources/certs/consul
	cfssl gencert \
		-ca=./resources/certs/ca.pem \
		-ca-key=./resources/certs/ca-key.pem \
		-config=./resources/cfssl/ca.json \
		-profile=default \
		./resources/cfssl/vault-csr.json | cfssljson -bare ./resources/certs/vault

secrets_create_gossip_key:
	@printf -- "$$(consul keygen)" > ./resources/gossip/key

secrets_init_k8s:
	@kubectl create secret -n $(K8S_NAMESPACE) generic consul \
		--from-literal="gossip-encryption-key=$$(cat ./resources/gossup/key)" \
		--from-file=resources/certs/ca.pem \
		--from-file=resources/certs/consul.pem \
		--from-file=resources/certs/consul-key.pem


#################################
# secrets initialisation in k8s #
#################################
#
# run this in the order in which they appear!
K8S_NAMESPACE=vault
# if you'd like to change the above values, create a file named
# `Makefile.properties` in this directory and override it
-include ./Makefile.properties

k8s_create_namespace:
	# creates the namespace for all the other operations to run in
	@kubectl create namespace $(K8S_NAMESPACE)

k8s_create_config:
	@kubectl create -n $(K8S_NAMESPACE) -f ./consul/configmap.yaml

k8s_create_service:
	@kubectl create -n $(K8S_NAMESPACE) -f ./consul/service.yaml

k8s_create_deployment:
	@kubectl create -n $(K8S_NAMESPACE) -f ./consul/statefulset.yaml

