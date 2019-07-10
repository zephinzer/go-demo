include ../../scripts/Makefile

define HELP_OUTPUT

direnv
^^^^^^
Directory specific environment variable loading

Link: https://github.com/direnv/direnv

endef

# installs terraform locally on the machine this recipe
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
	@curl https://api.github.com/repos/direnv/direnv/releases > ./.version_info

version_latest:
	@cat ./.version_info | jq '.[].tag_name' -r | sort -V | tail -n 1 | cut -f 2 -d 'v' > ./.version_latest
	@printf -- "latest version '$$(cat ./.version_latest)' stored in ./.version_latest\n"

download:
	@mkdir -p ./bin/$$(cat ./.version_latest)
	@curl -Lo "./bin/$$(cat ./.version_latest)/direnv-$(SYS_OS)-$(SYS_ARCH)" \
		"https://github.com/direnv/direnv/releases/download/v$$(cat ./.version_latest)/direnv.$(SYS_OS)-$(SYS_ARCH)"
	@chmod +x "./bin/$$(cat ./.version_latest)/direnv-$(SYS_OS)-$(SYS_ARCH)"

install:
	# if the `mkdir` below fails, run it manually with `sudo` and run
	# `chown $(whoami):$(whoami) /opt/terraform before running `make install` again
	@mkdir -p /opt/direnv
	@if ! [ -f "/opt/direnv/direnv-$$(cat ./.version_latest)$(BIN_EXT)" ]; then \
		cp "./bin/$$(cat ./.version_latest)/direnv-$(SYS_OS)-$(SYS_ARCH)$(BIN_EXT)" \
			"/opt/direnv/direnv-$$(cat ./.version_latest)$(BIN_EXT)"; \
	fi
	@if [ -f "/usr/bin/direnv$(BIN_EXT)" ]; then \
		ls -al "/usr/bin/direnv$(BIN_EXT)"; \
		printf -- "it seems like '/usr/bin/direnv$(BIN_EXT)' already exists, remove it before retrying.\n"; \
		exit 1; \
	fi
	@printf -- "requesting sudo access to create symlink at '/usr/bin/direnv$(BIN_EXT)'... "
	@sudo ln -s \
		"/opt/direnv/direnv-$$(cat ./.version_latest)$(BIN_EXT)" \
		"/usr/bin/direnv$(BIN_EXT)"

verify:
	@direnv version
	@ls -al $$(which direnv)
check: verify
