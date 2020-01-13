include ../../scripts/Makefile

define HELP_OUTPUT

AWS IAM Authenticator
^^^^^^^^^^^^^^^^^^^^^
Authentication tool to connect to a Kubernetes cluster via AWS IAM

Link: https://github.com/kubernetes-sigs/aws-iam-authenticator

endef

tldr_install:
	#
	# get latest versions
	@$(MAKE) version_info
	#
	# get latest version
	@$(MAKE) version_latest
	#
	# bring in the latest version locally
	@$(MAKE) download
	#
	# install it to our path
	@$(MAKE) install

version_info:
	@curl https://api.github.com/repos/kubernetes-sigs/aws-iam-authenticator/releases > ./.version_info
	
version_latest:
	@printf -- "$$(cat ./.version_info | jq '.[].tag_name' -r | cut -f 2 -d '-' | sort -V | tail -n 1 | tr -d 'v')" > ./.version_latest
	@printf -- "latest version '$$(cat ./.version_latest)' stored in ./.version_latest\n"

download:
	@mkdir -p ./bin/$$(cat ./.version_latest)
	@curl -Lo "./bin/$$(cat ./.version_latest)/aws-iam-authenticator-$(SYS_OS)-$(SYS_ARCH)$(BIN_EXT)" \
		"https://github.com/kubernetes-sigs/aws-iam-authenticator/releases/download/v$$(cat ./.version_latest)/aws-iam-authenticator_$$(cat ./.version_latest)_$(SYS_OS)_$(SYS_ARCH)$(BIN_EXT)"
	@chmod +x "./bin/$$(cat ./.version_latest)/aws-iam-authenticator-$(SYS_OS)-$(SYS_ARCH)$(BIN_EXT)"

install:
	@mkdir -p /opt/aws-iam-authenticator
	@if ! [ -f "/opt/aws-iam-authenticator/aws-iam-authenticator-$$(cat ./.version_latest)$(BIN_EXT)" ]; then \
		cp "./bin/$$(cat ./.version_latest)/aws-iam-authenticator-$(SYS_OS)-$(SYS_ARCH)$(BIN_EXT)" \
			"/opt/aws-iam-authenticator/aws-iam-authenticator-$$(cat ./.version_latest)$(BIN_EXT)"; \
	fi
	@if [ -f "/usr/bin/aws-iam-authenticator$(BIN_EXT)" ]; then \
		ls -al "/usr/bin/aws-iam-authenticator$(BIN_EXT)"; \
		printf -- "it seems like '/usr/bin/aws-iam-authenticator$(BIN_EXT)' already exists, remove it before retrying.\n"; \
		exit 1; \
	fi
	@printf -- "requesting sudo access to create symlink at '/usr/bin/aws-iam-authenticator$(BIN_EXT)'... "
	@sudo ln -s \
		"/opt/aws-iam-authenticator/aws-iam-authenticator-$$(cat ./.version_latest)$(BIN_EXT)" \
		"/usr/bin/aws-iam-authenticator$(BIN_EXT)"

check: verify
verify:
	@which aws-iam-authenticator \
		&& ls -al $$(which aws-iam-authenticator) \
		&& aws-iam-authenticator version
