include ../../scripts/Makefile

define HELP_OUTPUT

Terraform
^^^^^^^^
Terraform enables you to safely and predictably create, change,
and improve infrastructure. It is an open source tool that
codifies APIs into declarative configuration files that can
be shared amongst team members, treated as code, edited,
reviewed, and versioned.

Link: https://github.com/hashicorp/terraform

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
	@curl https://api.github.com/repos/hashicorp/terraform/releases > ./.version_info

version_latest:
	@cat ./.version_info | jq '.[].tag_name' -r | sort -V | tail -n 1 | cut -f 2 -d 'v' > ./.version_latest
	@printf -- "latest version '$$(cat ./.version_latest)' stored in ./.version_latest\n"

download:
	@mkdir -p ./bin/$$(cat ./.version_latest)
	@curl -Lo "./bin/$$(cat ./.version_latest)/terraform-$(SYS_OS)-$(SYS_ARCH).zip" \
		"https://releases.hashicorp.com/terraform/$$(cat ./.version_latest)/terraform_$$(cat ./.version_latest)_$(SYS_OS)_$(SYS_ARCH).zip"
	@unzip -d "./bin/$$(cat ./.version_latest)/" "./bin/$$(cat ./.version_latest)/terraform-$(SYS_OS)-$(SYS_ARCH).zip" \
		&& chmod +x "./bin/$$(cat ./.version_latest)/terraform$(BIN_EXT)"

install:
	# if the `mkdir` below fails, run it manually with `sudo` and run
	# `chown $(whoami):$(whoami) /opt/terraform before running `make install` again
	@mkdir -p /opt/terraform
	@if ! [ -f "/opt/terraform/terraform-$$(cat ./.version_latest)$(BIN_EXT)" ]; then \
		cp "./bin/$$(cat ./.version_latest)/terraform$(BIN_EXT)" \
			"/opt/terraform/terraform-$$(cat ./.version_latest)$(BIN_EXT)"; \
	fi
	@if [ -f "/usr/bin/terraform$(BIN_EXT)" ]; then \
		ls -al "/usr/bin/terraform$(BIN_EXT)"; \
		printf -- "it seems like '/usr/bin/terraform$(BIN_EXT)' already exists, remove it before retrying.\n"; \
		exit 1; \
	fi
	@printf -- "requesting sudo access to create symlink at '/usr/bin/terraform$(BIN_EXT)'... "
	@sudo ln -s \
		"/opt/terraform/terraform-$$(cat ./.version_latest)$(BIN_EXT)" \
		"/usr/bin/terraform$(BIN_EXT)"

verify:
	@terraform version
	@ls -al $$(which terraform)
check: verify
