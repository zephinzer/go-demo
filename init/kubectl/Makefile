define HELP_OUTPUT

Kubectl
^^^^^^^
This tool is used for communicating with a k8s cluster

Link: https://github.com/kubernetes/kubectl

endef

include ../../scripts/Makefile

version_info:
	@curl https://api.github.com/repos/kubernetes/kubernetes/releases > ./.version_info
	
version_latest:
	@printf -- "$$(cat ./.version_info | jq '.[].tag_name' -r | egrep -v "(alpha|beta|rc)" | sort -V | tail -n 1)" > ./.version_latest
	@printf -- "latest version '$$(cat ./.version_latest)' stored in ./.version_latest\n"

download:
	@mkdir -p ./bin/$$(cat ./.version_latest)
	@curl -Lo "./bin/$$(cat ./.version_latest)/kubectl-$(SYS_OS)-$(SYS_ARCH)$(BIN_EXT)" \
		"https://storage.googleapis.com/kubernetes-release/release/$$(cat ./.version_latest)/bin/$(SYS_OS)/$(SYS_ARCH)/kubectl$(BIN_EXT)"
	@chmod +x "./bin/$$(cat ./.version_latest)/kubectl-$(SYS_OS)-$(SYS_ARCH)$(BIN_EXT)"

install:
	@mkdir -p /opt/kubectl
	@if ! [ -f "/opt/kubectl/kubectl-$$(cat ./.version_latest)$(BIN_EXT)" ]; then \
		cp "./bin/$$(cat ./.version_latest)/kubectl-$(SYS_OS)-$(SYS_ARCH)$(BIN_EXT)" \
			"/opt/kubectl/kubectl-$$(cat ./.version_latest)$(BIN_EXT)"; \
	fi
	@if [ -f "/usr/bin/kubectl$(BIN_EXT)" ]; then \
		ls -al "/usr/bin/kubectl$(BIN_EXT)"; \
		printf -- "it seems like '/usr/bin/kubectl$(BIN_EXT)' already exists, remove it before retrying.\n"; \
		exit 1; \
	fi
	@printf -- "requesting sudo access to create symlink at '/usr/bin/kubectl$(BIN_EXT)'... "
	@sudo ln -s \
		"/opt/kubectl/kubectl-$$(cat ./.version_latest)$(BIN_EXT)" \
		"/usr/bin/kubectl$(BIN_EXT)"

check: verify
verify:
	@which kubectl \
		&& ls -al $$(which kubectl) \
		&& kubectl version --client --short
