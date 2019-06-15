include ../../scripts/Makefile

define HELP_OUTPUT

Helm
^^^^
Helm is a package manager for Kubernetes which helps us
spin up support services quickly and lets us rollback stuff

Link: https://github.com/helm/helm

endef

# installs the linkerd cli locally on the machine this recipe
# is being run on
# * intended for use on a fresh cluster
tldr_install:
	#
	# get latest versions
	@$(MAKE) version_info
	#
	# get latest version
	@$(MAKE) version_latest
	#
	# bring in the latest version
	@$(MAKE) download
	#
	# install it to our path
	@$(MAKE) install

# installs linkerd into the default cluster as defined
# by your kubectl config (run `kubectl config view` to see it)
tldr_k8s:
	@$(MAKE) k8s_pre_install_check
	@$(MAKE) k8s_install
	@$(MAKE) k8s_post_install_check

version_info:
	@curl https://api.github.com/repos/linkerd/linkerd2/releases > ./.version_info
	
version_latest:
	@printf -- "stable-$$(cat ./.version_info | jq '.[].tag_name' -r | grep 'stable-' | cut -f 2 -d '-' | sort -V | tail -n 1)" > ./.version_latest
	@printf -- "latest version '$$(cat ./.version_latest)' stored in ./.version_latest\n"

download:
	@mkdir -p ./bin/$$(cat ./.version_latest)
	@curl -Lo "./bin/$$(cat ./.version_latest)/linkerd-$(SYS_OS)$(BIN_EXT)" \
		"https://github.com/linkerd/linkerd2/releases/download/$$(cat ./.version_latest)/linkerd2-cli-$$(cat ./.version_latest)-$(SYS_OS)$(BIN_EXT)"
	@curl -Lo "./bin/$$(cat ./.version_latest)/linkerd-$(SYS_OS)$(BIN_EXT).sha256" \
		"https://github.com/linkerd/linkerd2/releases/download/$$(cat ./.version_latest)/linkerd2-cli-$$(cat ./.version_latest)-$(SYS_OS)$(BIN_EXT).sha256"
	@chmod +x "./bin/$$(cat ./.version_latest)/linkerd-$(SYS_OS)$(BIN_EXT)"

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
