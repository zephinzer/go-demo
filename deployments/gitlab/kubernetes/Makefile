-include ../../../scripts/Makefile

define HELP_OUTPUT

GitLab on Kubernetes
^^^^^^^^^^^^^^^^^^^^

Run `make minikube` to create a new Minikube profile
named `gitlab`.

endef
export HELP_OUTPUT

minikube: check_prereq_minikube
	@if ! minikube status --profile gitlab; then \
		minikube start \
			--cpus 4 \
			--disk-size 16g \
			--memory 6144 \
			--profile gitlab \
			--vm-driver kvm2 \
			-v=9; \
	else \
		printf -- 'minikube already started\n'; \
	fi

kubernetes:
	@kubectl apply -f ./helm.yaml

kubernetes_destroy:
	@kubectl delete -f ./helm.yaml
