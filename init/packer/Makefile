include ../../scripts/Makefile

define HELP_OUTPUT

Packer
^^^^^^
Packer is a tool for creating identical machine images
for multiple platforms from a single source configuration.

Link: https://github.com/hashicorp/packer

endef

# installs packer locally on the machine this recipe
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


# note: this does things differently from the rest because hashicorp
#       releases based on tags, if we use the `release` path, we'll
#       only get the "nightly" release
version_info:
	@curl "https://api.github.com/repos/hashicorp/packer/tags" > ./.version_info

version_latest:
	@cat ./.version_info | jq '.[].name' -r | egrep -v "(alpha|beta|rc)" | sort -V | tail -n 1 | sed -e 's|v||g' > ./.version_latest
	@printf -- "latest version '$$(cat ./.version_latest)' stored in ./.version_latest\n"

download:
	@mkdir -p ./bin/$$(cat ./.version_latest)
	@curl -Lo "./bin/$$(cat ./.version_latest)/packer-$(SYS_OS)-$(SYS_ARCH).zip" \
		"https://releases.hashicorp.com/packer/$$(cat ./.version_latest)/packer_$$(cat ./.version_latest)_$(SYS_OS)_$(SYS_ARCH).zip"
	@unzip -d "./bin/$$(cat ./.version_latest)/" "./bin/$$(cat ./.version_latest)/packer-$(SYS_OS)-$(SYS_ARCH).zip" \
		&& chmod +x "./bin/$$(cat ./.version_latest)/packer$(BIN_EXT)"

install:
	@mkdir -p /opt/packer
	@if ! [ -f "/opt/packer/packer-$$(cat ./.version_latest)$(BIN_EXT)" ]; then \
		cp "./bin/$$(cat ./.version_latest)/packer$(BIN_EXT)" \
			"/opt/packer/packer-$$(cat ./.version_latest)$(BIN_EXT)"; \
	fi
	@if [ -f "/usr/bin/packer$(BIN_EXT)" ]; then \
		ls -al "/usr/bin/packer$(BIN_EXT)"; \
		printf -- "it seems like '/usr/bin/packer$(BIN_EXT)' already exists, remove it before retrying.\n"; \
		exit 1; \
	fi
	@printf -- "requesting sudo access to create symlink at '/usr/bin/packer$(BIN_EXT)'... "
	@sudo ln -s \
		"/opt/packer/packer-$$(cat ./.version_latest)$(BIN_EXT)" \
		"/usr/bin/packer$(BIN_EXT)"

verify:
	@packer version
	@ls -al $$(which packer)
check: verify
