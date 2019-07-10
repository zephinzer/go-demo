include ../../scripts/Makefile

define HELP_OUTPUT

Consul by Hashicorp
^^^^^^^^^^^^^^^^^^^
Consul is a distributed, highly available, and data center aware solution to connect and configure applications across dynamic, distributed infrastructure. 

Link: https://github.com/hashicorp/consul

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
	@curl "https://api.github.com/repos/hashicorp/consul/tags" > ./.version_info

version_latest:
	@cat ./.version_info | jq '.[].name' -r | egrep -v "(alpha|beta|rc)" | sort -V | tail -n 1 | sed -e 's|v||g' > ./.version_latest
	@printf -- "latest version '$$(cat ./.version_latest)' stored in ./.version_latest\n"

download:
	@mkdir -p ./bin/$$(cat ./.version_latest)
	@curl -Lo "./bin/$$(cat ./.version_latest)/consul-$(SYS_OS)-$(SYS_ARCH).zip" \
		"https://releases.hashicorp.com/consul/$$(cat ./.version_latest)/consul_$$(cat ./.version_latest)_$(SYS_OS)_$(SYS_ARCH).zip"
	@unzip -d "./bin/$$(cat ./.version_latest)/" "./bin/$$(cat ./.version_latest)/consul-$(SYS_OS)-$(SYS_ARCH).zip" \
		&& chmod +x "./bin/$$(cat ./.version_latest)/consul$(BIN_EXT)"

install:
	# if the `mkdir` below fails, run it manually with `sudo` and run
	# `chown $(whoami):$(whoami) /opt/consul before running `make install` again
	@mkdir -p /opt/consul
	@if ! [ -f "/opt/consul/consul-$$(cat ./.version_latest)$(BIN_EXT)" ]; then \
		cp "./bin/$$(cat ./.version_latest)/consul$(BIN_EXT)" \
			"/opt/consul/consul-$$(cat ./.version_latest)$(BIN_EXT)"; \
	fi
	@if [ -f "/usr/bin/consul$(BIN_EXT)" ]; then \
		ls -al "/usr/bin/consul$(BIN_EXT)"; \
		printf -- "it seems like '/usr/bin/consul$(BIN_EXT)' already exists, remove it before retrying.\n"; \
		exit 1; \
	fi
	@printf -- "requesting sudo access to create symlink at '/usr/bin/consul$(BIN_EXT)'... "
	@sudo ln -s \
		"/opt/consul/consul-$$(cat ./.version_latest)$(BIN_EXT)" \
		"/usr/bin/consul$(BIN_EXT)"

check: verify
verify:
	@consul version
	@ls -al $$(which consul)
