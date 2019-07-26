-include ../../../scripts/Makefile

define HELP_OUTPUT

GitLab on Kubernetes
^^^^^^^^^^^^^^^^^^^^

Run `make minikube` to create a new Minikube profile
named `gitlab`.

endef
export HELP_OUTPUT

certs_create:
	@mkdir -p ./.helm
	#
	# ca
	#
	@openssl genrsa -out ./.helm/ca.key.pem 4096
	@openssl req -new -x509 -sha256 \
		-key ./.helm/ca.key.pem \
		-days 3650 \
		-extensions v3_ca \
		-subj "/C=SG/ST=Singapore/L=Singapore/O=zephinzer-demo/OU=helmca/CN=helmca" \
		-out ./.helm/ca.cert.pem
	#
	# helm
	#
	@openssl genrsa -out ./.helm/helm.key.pem 4096
	@openssl req -new -sha256 \
		-key ./.helm/helm.key.pem \
		-subj "/C=SG/ST=Singapore/L=Singapore/O=zephinzer-demo/OU=helm/CN=helmuser" \
		-out ./.helm/helm.csr.pem
	@openssl x509 -req \
		-days 3650 \
		-CAcreateserial \
		-CA ./.helm/ca.cert.pem \
		-CAkey ./.helm/ca.key.pem \
		-in ./.helm/helm.csr.pem \
		-out ./.helm/helm.cert.pem
	#
	# tiller
	#
	@openssl genrsa -out ./.helm/tiller.key.pem 4096
	@openssl req -new -sha256 \
		-key ./.helm/tiller.key.pem \
		-subj "/C=SG/ST=Singapore/L=Singapore/O=zephinzer-demo/OU=tiller/CN=tilleruser" \
		-out ./.helm/tiller.csr.pem
	@openssl x509 -req \
		-days 3650 \
		-CAcreateserial \
		-CA ./.helm/ca.cert.pem \
		-CAkey ./.helm/ca.key.pem \
		-in ./.helm/tiller.csr.pem \
		-out ./.helm/tiller.cert.pem

certs_destroy:
	@rm -rf ./.helm/*


TILLER_NAMESPACE=gitlab
helm_create:
	@helm init \
		--service-account tiller \
		--tiller-namespace $(TILLER_NAMESPACE)
		# --override 'spec.template.spec.containers[0].command'='{/tiller,--storage=secret}' \
		# --history-max 500 \
		# --tiller-tls \
		# --tiller-tls-cert ./.helm/tiller.cert.pem \
		# --tiller-tls-key ./.helm/tiller.key.pem \
		# --tiller-tls-verify \
		# --tls-ca-cert ./.helm/ca.cert.pem \

helm_version:
	@helm version

helm_destroy:
	@kubectl delete -n $(TILLER_NAMESPACE) service tiller-deploy
	@kubectl delete -n $(TILLER_NAMESPACE) deployment tiller-deploy
	@kubectl delete -n $(TILLER_NAMESPACE) secret tiller-secret

	