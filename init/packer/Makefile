include ../../scripts/Makefile

define HELP_OUTPUT

Packer by Hashicorp
^^^^^^^^^^^^^^^^^^^
Packer is a tool used to create machine images for use
in cloud providers.

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
	@curl -Lo "./bin/$$(cat ./.version_latest)/packer-$(SYS_OS)-$(SYS_ARCH)$(BIN_EXT)" \
		"https://releases.hashicorp.com/packer/$$(cat ./.version_latest)/packer_$$(cat ./.version_latest)_$(SYS_OS)_$(SYS_ARCH).zip"
	@chmod +x "./bin/$$(cat ./.version_latest)/packer-$(SYS_OS)-$(SYS_ARCH)$(BIN_EXT)"

# install:
# 	@mkdir -p /opt/minikube
# 	@if ! [ -f "/opt/minikube/minikube-$$(cat ./.version_latest)$(BIN_EXT)" ]; then \
# 		cp "./bin/$$(cat ./.version_latest)/minikube-$(SYS_OS)-$(SYS_ARCH)$(BIN_EXT)" \
# 			"/opt/minikube/minikube-$$(cat ./.version_latest)$(BIN_EXT)"; \
# 	fi
# 	@if [ -f "/usr/bin/minikube$(BIN_EXT)" ]; then \
# 		ls -al "/usr/bin/minikube$(BIN_EXT)"; \
# 		printf -- "it seems like '/usr/bin/minikube$(BIN_EXT)' already exists, remove it before retrying.\n"; \
# 		exit 1; \
# 	fi
# 	@printf -- "requesting sudo access to create symlink at '/usr/bin/minikube$(BIN_EXT)'... "
# 	@sudo ln -s \
# 		"/opt/minikube/minikube-$$(cat ./.version_latest)$(BIN_EXT)" \
# 		"/usr/bin/minikube$(BIN_EXT)"

# verify:
# 	@minikube version
# 	@ls -al $$(which minikube)
# check: verify
