-include ../../../scripts/Makefile

define HELP_OUTPUT

GitLab on Kubernetes
^^^^^^^^^^^^^^^^^^^^

# Get Started
Run `make certs_create` to generate a new set of certs.

Run `make helm_init` to initialise helm.

# Required Environment
Use `direnv allow .` to enable usage of the `.envrc`.
Alternatively, run `source ./.envrc` to load the environemnt
variables **after creating the certs**.

endef
export HELP_OUTPUT

certs_create:
	@mkdir -p ./.ssl
	$(MAKE) create_ca_cert
	$(MAKE) create_user_cert USE_CASE=helm
	$(MAKE) create_user_cert USE_CASE=tiller
certs_destroy:
	@rm -rf ./.ssl/*

CERT_DIR=./.ssl
CA_CERT_NAME=ca
create_ca_cert:
	@openssl genrsa -out $(CERT_DIR)/$(CA_CERT_NAME).key.pem 4096
	@openssl req -new -x509 \
		-key $(CERT_DIR)/$(CA_CERT_NAME).key.pem \
		-days 3650 \
		-extensions v3_ca \
		-subj "/C=SG/ST=Singapore/L=Singapore/O=zephinzer-demo/OU=ca/CN=ca" \
		-out $(CERT_DIR)/$(CA_CERT_NAME).cert.pem
create_user_cert:
	@openssl genrsa -out $(CERT_DIR)/${USE_CASE}.key.pem 4096
	@openssl req -new -sha256 \
		-key $(CERT_DIR)/${USE_CASE}.key.pem \
		-subj "/C=SG/ST=Singapore/L=Singapore/O=zephinzer-demo/OU=${USE_CASE}/CN=${USE_CASE}" \
		-out $(CERT_DIR)/${USE_CASE}.csr.pem
	@openssl x509 -req \
		-days 3650 \
		-CAcreateserial \
		-CA $(CERT_DIR)/$(CA_CERT_NAME).cert.pem \
		-CAkey $(CERT_DIR)/$(CA_CERT_NAME).key.pem \
		-in $(CERT_DIR)/${USE_CASE}.csr.pem \
		-out $(CERT_DIR)/${USE_CASE}.cert.pem
	@rm -rf $(CERT_DIR)/${USE_CASE}.csr.pem

TILLER_NAMESPACE=gitlab
helm_init:
	@helm init \
		--tiller-tls \
		--tiller-tls-cert $(CERT_DIR)/tiller.cert.pem \
		--tiller-tls-key $(CERT_DIR)/tiller.key.pem \
		--tiller-tls-verify \
		--tls-ca-cert $(CERT_DIR)/ca.cert.pem \
		--tiller-tls-hostname tiller \
		--history-max 500 \
		--override 'spec.template.spec.containers[0].command'='{/tiller,--storage=secret}' \
		--upgrade \
		--force-upgrade \
		--wait \
		--debug \
		--service-account tiller \
		--tiller-namespace $(TILLER_NAMESPACE)
helm_destroy:
	-@kubectl delete -n $(TILLER_NAMESPACE) service tiller-deploy
	-@kubectl delete -n $(TILLER_NAMESPACE) deployment tiller-deploy
	-@kubectl delete -n $(TILLER_NAMESPACE) secret tiller-secret
helm_version:
	@helm version

gitlab_get:
	@helm repo add gitlab https://charts.gitlab.io
	@helm repo update

gitlab_install:
	@helm install \
		--namespace gitlab \
		--values ./values.yaml \
		gitlab/gitlab
