-include ../../../scripts/Makefile

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
	@kubectl apply -f ./k8s-prereqs.yaml

kubernetes_reset:
	-@kubectl delete -f ./k8s-prereqs.yaml
	$(MAKE) kubernetes
