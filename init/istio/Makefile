include ../../scripts/Makefile

# installs the client side stuff
tldr_install:
	@$(MAKE) version_info
	@$(MAKE) version_latest
	@$(MAKE) download
	@$(MAKE) install

# installs the server side stuff
tldr_k8s_install:
	@$(MAKE) k8s_preinstall
	@$(MAKE) k8s_install_init
	@printf -- "waiting for about 30 seconds before throwing in istio (otherwise it may fail...)"
	@sleep 1 && printf '.' && sleep 1 && printf '.' && sleep 1 && printf '.' && sleep 1 && printf '.' && sleep 1 && printf '.'
	@sleep 1 && printf '.' && sleep 1 && printf '.' && sleep 1 && printf '.' && sleep 1 && printf '.' && sleep 1 && printf '.'
	@sleep 1 && printf '.' && sleep 1 && printf '.' && sleep 1 && printf '.' && sleep 1 && printf '.' && sleep 1 && printf '.'
	@sleep 1 && printf '.' && sleep 1 && printf '.' && sleep 1 && printf '.' && sleep 1 && printf '.' && sleep 1 && printf '.'
	@sleep 1 && printf '.' && sleep 1 && printf '.' && sleep 1 && printf '.' && sleep 1 && printf '.' && sleep 1 && printf '.'
	@sleep 1 && printf '.' && sleep 1 && printf '.' && sleep 1 && printf '.' && sleep 1 && printf '.' && sleep 1 && printf '.\n'
	@$(MAKE) k8s_install_istio

# uninstalls the server side stuff
tldr_k8s_purge:
	@$(MAKE) k8s_purge_istio
	@$(MAKE) k8s_purge_init
	@$(MAKE) k8s_postpurge

##########
# client #
##########

version_info:
	@curl https://api.github.com/repos/istio/istio/releases > ./.version_info
version_latest:
	cat ./.version_info | jq '.[].tag_name' -r | grep -v rc | grep -v snapshot | sort -V | tail -n 1 > ./.version_latest
	@printf -- "latest version '$$(cat ./.version_latest)' stored in ./.version_latest\n"

download:
	@mkdir -p ./bin/$$(cat ./.version_latest)
	@curl -Lo "./bin/$$(cat ./.version_latest)/istio-$(SYS_OS_TRUNC).tar.gz" \
		"https://github.com/istio/istio/releases/download/$$(cat ./.version_latest)/istio-$$(cat ./.version_latest)-$(SYS_OS_TRUNC).tar.gz"
	@curl -Lo "./bin/$$(cat ./.version_latest)/istio-$(SYS_OS_TRUNC).tar.gz.sha256" \
		"https://github.com/istio/istio/releases/download/$$(cat ./.version_latest)/istio-$$(cat ./.version_latest)-$(SYS_OS_TRUNC).tar.gz.sha256"
	@tar -C ./bin/$$(cat ./.version_latest) -xzvf "./bin/$$(cat ./.version_latest)/istio-$(SYS_OS_TRUNC).tar.gz" \
		&& mv ./bin/$$(cat ./.version_latest)/istio-$$(cat ./.version_latest)/* ./bin/$$(cat ./.version_latest) \
		&& rm -rf ./bin/$$(cat ./.version_latest)/istio-$$(cat ./.version_latest)

install:
	@mkdir -p /opt/istioctl
	@if ! [ -f "/opt/istioctl/istioctl-$$(cat ./.version_latest)$(BIN_EXT)" ]; then \
		cp "./bin/$$(cat ./.version_latest)/bin/istioctl$(BIN_EXT)" \
			"/opt/istioctl/istioctl-$$(cat ./.version_latest)$(BIN_EXT)"; \
	fi
	@if [ -f "/usr/bin/istioctl$(BIN_EXT)" ]; then \
		ls -al "/usr/bin/istioctl$(BIN_EXT)"; \
		printf -- "it seems like '/usr/bin/istioctl$(BIN_EXT)' already exists, remove it before retrying.\n"; \
		exit 1; \
	fi
	@printf -- "requesting sudo access to create symlink at '/usr/bin/istioctl$(BIN_EXT)'... "
	@sudo ln -s \
		"/opt/istioctl/istioctl-$$(cat ./.version_latest)$(BIN_EXT)" \
		"/usr/bin/istioctl$(BIN_EXT)"

#################
# on kubernetes #
#################

HELM_OPT := --tls --tiller-namespace=tiller

k8s_preinstall:
	@kubectl apply -f ./resources/namespace.yaml
	@kubectl apply -f ./resources/clusterrole.yaml

k8s_postpurge:
	-@kubectl delete -f ./resources/clusterrole.yaml
	-@kubectl delete -f ./resources/namespace.yaml

k8s_install_init:
	@helm install ./bin/$$(cat ./.version_latest)/install/kubernetes/helm/istio-init $(HELM_OPT) \
		--name istio-init \
		--namespace istio-system
k8s_purge_init:
	-@helm delete $(HELM_OPT) \
		--purge istio-init

k8s_install_istio:
	@printf -- 'NOTE: if you ran this immediately after running the istio-init, this may fail, wait for awhile and try again.\n'
	@helm install ./bin/$$(cat ./.version_latest)/install/kubernetes/helm/istio $(HELM_OPT) \
		-f ./resources/values.yaml \
		--name istio \
		--namespace istio-system
k8s_update_istio:
	@helm upgrade $(HELM_OPT) \
		-f ./resources/values.yaml \
		istio ./bin/$$(cat ./.version_latest)/install/kubernetes/helm/istio
k8s_purge_istio:
	-@helm delete $(HELM_OPT) \
		--purge istio

k8s_check:
	@kubectl get services -n istio-system -o wide
	@kubectl get pods -n istio-system -o wide

k8s_grafana:
	@kubectl port-forward -n istio-system $$(kubectl get pods -n istio-system | grep grafana | cut -f 1 -d ' ') $$(kubectl describe service -n istio-system grafana | grep TargetPort | cut -f 2- -d ' ' | sed -e 's| ||g' | cut -f 1 -d '/'):$$(kubectl describe service -n istio-system grafana | grep TargetPort | cut -f 2- -d ' ' | sed -e 's| ||g' | cut -f 1 -d '/')

k8s_jaegar:
	@kubectl port-forward -n istio-system $$(kubectl get pods -n istio-system | grep tracing | cut -f 1 -d ' ') $$(kubectl describe service -n istio-system jaeger-query | grep TargetPort | cut -f 2- -d ' ' | sed -e 's| ||g' | cut -f 1 -d '/'):$$(kubectl describe service -n istio-system jaeger-query | grep TargetPort | cut -f 2- -d ' ' | sed -e 's| ||g' | cut -f 1 -d '/')

k8s_kiali:
	@kubectl port-forward -n istio-system $$(kubectl get pods -n istio-system | grep kiali | cut -f 1 -d ' ') $$(kubectl describe service -n istio-system kiali | grep TargetPort | cut -f 2- -d ' ' | sed -e 's| ||g' | cut -f 1 -d '/'):$$(kubectl describe service -n istio-system kiali | grep TargetPort | cut -f 2- -d ' ' | sed -e 's| ||g' | cut -f 1 -d '/')

k8s_enable:
	@if [ "${NS}" = "" ]; then \
		printf -- "specify the \$$NS variable to proceed (eg. 'make k8s_enable NS=my-namespace').\n"; \
		exit 1; \
	fi
	@kubectl label namespace ${NS} istio-injection=enabled 
