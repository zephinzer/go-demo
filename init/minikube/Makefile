include ../../scripts/Makefile

define HELP_OUTPUT

Minikube
^^^^^^^^
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


version_info:
	@curl https://api.github.com/repos/kubernetes/minikube/releases > ./.version_info

version_latest:
	@cat ./.version_info | jq '.[].tag_name' -r | sort -V | tail -n 1 > ./.version_latest
	@printf -- "latest version '$$(cat ./.version_latest)' stored in ./.version_latest\n"

download:
	@mkdir -p ./bin/$$(cat ./.version_latest)
	@curl -Lo "./bin/$$(cat ./.version_latest)/minikube-$(SYS_OS)-$(SYS_ARCH)$(BIN_EXT)" \
		"https://storage.googleapis.com/minikube/releases/$$(cat ./.version_latest)/minikube-$(SYS_OS)-$(SYS_ARCH)$(BIN_EXT)"
	@curl -Lo "./bin/$$(cat ./.version_latest)/minikube-$(SYS_OS)-$(SYS_ARCH)$(BIN_EXT).sha256" \
		"https://storage.googleapis.com/minikube/releases/$$(cat ./.version_latest)/minikube-$(SYS_OS)-$(SYS_ARCH)$(BIN_EXT).sha256"
	@chmod +x "./bin/$$(cat ./.version_latest)/minikube-$(SYS_OS)-$(SYS_ARCH)$(BIN_EXT)"

install:
	@mkdir -p /opt/minikube
	@if ! [ -f "/opt/minikube/minikube-$$(cat ./.version_latest)$(BIN_EXT)" ]; then \
		cp "./bin/$$(cat ./.version_latest)/minikube-$(SYS_OS)-$(SYS_ARCH)$(BIN_EXT)" \
			"/opt/minikube/minikube-$$(cat ./.version_latest)$(BIN_EXT)"; \
	fi
	@if [ -f "/usr/bin/minikube$(BIN_EXT)" ]; then \
		ls -al "/usr/bin/minikube$(BIN_EXT)"; \
		printf -- "it seems like '/usr/bin/minikube$(BIN_EXT)' already exists, remove it before retrying.\n"; \
		exit 1; \
	fi
	@printf -- "requesting sudo access to create symlink at '/usr/bin/minikube$(BIN_EXT)'... "
	@sudo ln -s \
		"/opt/minikube/minikube-$$(cat ./.version_latest)$(BIN_EXT)" \
		"/usr/bin/minikube$(BIN_EXT)"

verify:
	@minikube version
	@ls -al $$(which minikube)
check: verify

k8s_status:
	minikube status
