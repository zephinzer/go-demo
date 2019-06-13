include ../Makefile

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
	@cd ./bin/$$(cat ./.version_latest) \
		&& tar -xzvf "./istio-$(SYS_OS_TRUNC).tar.gz" \
		&& mv ./bin/$$(cat ../../.version_latest)/istio-$$(cat ../../.version_latest)/* ./bin/$$(cat ../../.version_latest) \
		&& rm -rf ./bin/$$(cat ../../.version_latest)/istio-$$(cat ../../.version_latest)

install:
	@mkdir -p /opt/linkerd
	@if ! [ -f "/opt/linkerd/linkerd-$$(cat ./.version_latest)$(BIN_EXT)" ]; then \
		cp "./bin/$$(cat ./.version_latest)/linkerd-$(SYS_OS)$(BIN_EXT)" \
			"/opt/linkerd/linkerd-$$(cat ./.version_latest)$(BIN_EXT)"; \
	fi
	@if [ -f "/usr/bin/linkerd$(BIN_EXT)" ]; then \
		ls -al "/usr/bin/linkerd$(BIN_EXT)"; \
		printf -- "it seems like '/usr/bin/linkerd$(BIN_EXT)' already exists, remove it before retrying.\n"; \
		exit 1; \
	fi
	@printf -- "requesting sudo access to create symlink at '/usr/bin/linkerd$(BIN_EXT)'... "
	@sudo ln -s \
		"/opt/linkerd/linkerd-$$(cat ./.version_latest)$(BIN_EXT)" \
		"/usr/bin/linkerd$(BIN_EXT)"

check: verify
verify:
	@linkerd version && ls -al $$(which linkerd)

k8s_pre_install_check:
	@linkerd check --pre

k8s_install:
	@linkerd install | kubectl apply -f -

k8s_post_install_check:
	@linkerd check

k8s_get_pods:
	@kubectl get pods -n linkerd -o wide

linkerd_dashboard:
	@linkerd dashboard
