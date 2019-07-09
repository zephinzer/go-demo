include ../../scripts/Makefile

define HELP_OUTPUT

Kustomize
^^^^^^^^
Kustomize lets you customize raw, template-free YAML files for
multiple purposes, leaving the original YAML untouched and
usable as is.

Link: https://github.com/kubernetes-sigs/kustomize

endef

# installs kustomize locally on the machine this recipe
# is being run on
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
	@curl https://api.github.com/repos/kubernetes-sigs/kustomize/releases > ./.version_info

version_latest:
	@cat ./.version_info | jq '.[].tag_name' -r | sort -V | tail -n 1 | cut -f 2 -d 'v' > ./.version_latest
	@printf -- "latest version '$$(cat ./.version_latest)' stored in ./.version_latest\n"

download:
	@mkdir -p ./bin/$$(cat ./.version_latest)
	@curl -Lo "./bin/$$(cat ./.version_latest)/kustomize-$(SYS_OS)-$(SYS_ARCH)$(BIN_EXT)" \
		"https://github.com/kubernetes-sigs/kustomize/releases/download/v$$(cat ./.version_latest)/kustomize_$$(cat ./.version_latest)_$(SYS_OS)_$(SYS_ARCH)$(BIN_EXT)"
	@chmod +x "./bin/$$(cat ./.version_latest)/kustomize-$(SYS_OS)-$(SYS_ARCH)$(BIN_EXT)"

install:
	@mkdir -p /opt/kustomize
	@if ! [ -f "/opt/kustomize/kustomize-$$(cat ./.version_latest)$(BIN_EXT)" ]; then \
		cp "./bin/$$(cat ./.version_latest)/kustomize-$(SYS_OS)-$(SYS_ARCH)$(BIN_EXT)" \
			"/opt/kustomize/kustomize-$$(cat ./.version_latest)$(BIN_EXT)"; \
	fi
	@if [ -f "/usr/bin/kustomize$(BIN_EXT)" ]; then \
		ls -al "/usr/bin/kustomize$(BIN_EXT)"; \
		printf -- "it seems like '/usr/bin/kustomize$(BIN_EXT)' already exists, remove it before retrying.\n"; \
		exit 1; \
	fi
	@printf -- "requesting sudo access to create symlink at '/usr/bin/kustomize$(BIN_EXT)'... "
	@sudo ln -s \
		"/opt/kustomize/kustomize-$$(cat ./.version_latest)$(BIN_EXT)" \
		"/usr/bin/kustomize$(BIN_EXT)"

verify:
	@kustomize version
	@ls -al $$(which kustomize)
check: verify
