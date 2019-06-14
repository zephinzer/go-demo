include ../../scripts/Makefile

# this assumes you installed helm/tiller via the method
# found at ~/init/helm directory
HELM_TLS_OPTS := --tls --tiller-namespace=tiller

prerequisites:
	@which helm \
		&& ls -al $$(which helm) \
		&& helm version $(HELM_TLS_OPTS)

install: prerequisites
	helm install $(HELM_TLS_OPTS) stable/prometheus-operator

download: prerequisites
	helm get $(HELM_TLS_OPTS) stable/prometheus-operator
