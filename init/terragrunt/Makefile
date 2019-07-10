include ../../scripts/Makefile

define HELP_OUTPUT

Terragrunt
^^^^^^^^^^
Terragrunt is a thin wrapper for Terraform that provides extra
tools for working with multiple Terraform modules.

Link: https://github.com/gruntwork-io/terragrunt

endef

# installs terragrunt locally on the machine this recipe
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
	@curl https://api.github.com/repos/gruntwork-io/terragrunt/releases > ./.version_info

version_latest:
	@cat ./.version_info | jq '.[].tag_name' -r | sort -V | tail -n 1 | cut -f 2 -d 'v' > ./.version_latest
	@printf -- "latest version '$$(cat ./.version_latest)' stored in ./.version_latest\n"

download:
	@mkdir -p ./bin/$$(cat ./.version_latest)
	@curl -Lo "./bin/$$(cat ./.version_latest)/terragrunt-$(SYS_OS)-$(SYS_ARCH)$(BIN_EXT)" \
		"https://github.com/gruntwork-io/terragrunt/releases/download/v$$(cat ./.version_latest)/terragrunt_$(SYS_OS)_$(SYS_ARCH)"
	@chmod +x "./bin/$$(cat ./.version_latest)/terragrunt-$(SYS_OS)-$(SYS_ARCH)$(BIN_EXT)"

install:
	# if the `mkdir` below fails, run it manually with `sudo` and run
	# `chown $(whoami):$(whoami) /opt/terragrunt before running `make install` again
	@mkdir -p /opt/terragrunt
	@if ! [ -f "/opt/terragrunt/terragrunt-$$(cat ./.version_latest)$(BIN_EXT)" ]; then \
		cp "./bin/$$(cat ./.version_latest)/terragrunt-$(SYS_OS)-$(SYS_ARCH)$(BIN_EXT)" \
			"/opt/terragrunt/terragrunt-$$(cat ./.version_latest)$(BIN_EXT)"; \
	fi
	@if [ -f "/usr/bin/terragrunt$(BIN_EXT)" ]; then \
		ls -al "/usr/bin/terragrunt$(BIN_EXT)"; \
		printf -- "it seems like '/usr/bin/terragrunt$(BIN_EXT)' already exists, remove it before retrying.\n"; \
		exit 1; \
	fi
	@printf -- "requesting sudo access to create symlink at '/usr/bin/terragrunt$(BIN_EXT)'... "
	@sudo ln -s \
		"/opt/terragrunt/terragrunt-$$(cat ./.version_latest)$(BIN_EXT)" \
		"/usr/bin/terragrunt$(BIN_EXT)"

verify:
	@terragrunt -v
	@ls -al $$(which terragrunt)
check: verify
