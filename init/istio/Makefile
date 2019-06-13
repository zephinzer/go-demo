include ../Makefile

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

k8s_preinstall:
	@kubectl apply -f ./resources/namespace.yaml
	@kubectl apply -f ./resources/clusterrole.yaml

k8s_postpurge:
	-@kubectl delete -f ./resources/clusterrole.yaml
	-@kubectl delete -f ./resources/namespace.yaml

k8s_install_init:
	@helm install ./bin/$$(cat ./.version_latest)/install/kubernetes/helm/istio-init \
		--tls \
		--tiller-namespace=tiller \
		--name istio-init \
		--namespace istio-system
k8s_purge_init:
	-@helm delete \
		--tls \
		--tiller-namespace=tiller \
		--purge istio-init

k8s_install_istio:
	@printf -- 'NOTE: if you ran this immediately after running the istio-init, this may fail, wait for awhile and try again.\n'
	@helm install ./bin/$$(cat ./.version_latest)/install/kubernetes/helm/istio \
		--tls \
		--tiller-namespace=tiller \
		--set global.proxy.resources.limits.cpu=100m \
		--set global.proxy.resources.limits.memory=50Mi \
		--set global.proxy.resources.requests.cpu=50m \
		--set global.proxy.resources.requests.memory=25Mi \
		--set grafana.enabled=true \
		--set mixer.telemetry.resources.limits.cpu=600m \
		--set mixer.telemetry.resources.limits.memory=700G \
		--set mixer.telemetry.resources.requests.cpu=300m \
		--set mixer.telemetry.resources.requests.memory=500Mi \
		--set pilot.resources.requests.cpu=200m \
		--set pilot.resources.requests.memory=500Mi \
		--set tracing.enabled=true \
		--set tracing.zipkin.resources.limits.cpu=150m \
		--set tracing.zipkin.resources.limits.memory=500Mi \
		--set tracing.zipkin.resources.requests.cpu=100m \
		--set tracing.zipkin.resources.requests.memory=250Mi \
		--name istio \
		--namespace istio-system
k8s_purge_istio:
	-@helm delete \
		--tls \
		--tiller-namespace=tiller \
		--purge istio

k8s_check:
	@kubectl get services -n istio-system -o wide
	@kubectl get pods -n istio-system -o wide

k8s_grafana:
	@kubectl port-forward -n istio-system $$(kubectl get pods -n istio-system | grep grafana | cut -f 1 -d ' ') $$(kubectl describe service -n istio-system grafana | grep TargetPort | cut -f 2- -d ' ' | sed -e 's| ||g' | cut -f 1 -d '/'):$$(kubectl describe service -n istio-system grafana | grep TargetPort | cut -f 2- -d ' ' | sed -e 's| ||g' | cut -f 1 -d '/')

k8s_jaegar:
	@kubectl port-forward -n istio-system $$(kubectl get pods -n istio-system | grep tracing | cut -f 1 -d ' ') $$(kubectl describe service -n istio-system jaeger-query | grep TargetPort | cut -f 2- -d ' ' | sed -e 's| ||g' | cut -f 1 -d '/'):$$(kubectl describe service -n istio-system jaeger-query | grep TargetPort | cut -f 2- -d ' ' | sed -e 's| ||g' | cut -f 1 -d '/')

k8s_enable:
	@if [ "${NS}" = "" ]; then \
		printf -- "specify the \$$NS variable to proceed (eg. 'make k8s_enable NS=my-namespace').\n"; \
		exit 1; \
	fi
	@kubectl label namespace ${NS} istio-injection=enabled 
