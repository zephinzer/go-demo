include ../../scripts/Makefile

version_info:
	@curl https://api.github.com/repos/kubernetes/kompose/releases > ./.version_info
	
version_latest:
	@printf -- "$$(cat ./.version_info | jq '.[].tag_name' -r | cut -f 2 -d '-' | sort -V | tail -n 1)" > ./.version_latest
	@printf -- "latest version '$$(cat ./.version_latest)' stored in ./.version_latest\n"

download:
	@mkdir -p ./bin/$$(cat ./.version_latest)
	@curl -Lo "./bin/$$(cat ./.version_latest)/kompose-$(SYS_OS)-$(SYS_ARCH)$(BIN_EXT)" \
		"https://github.com/kubernetes/kompose/releases/download/$$(cat ./.version_latest)/kompose-$(SYS_OS)-$(SYS_ARCH)"
	@chmod +x "./bin/$$(cat ./.version_latest)/kompose-$(SYS_OS)-$(SYS_ARCH)$(BIN_EXT)"

install:
	@mkdir -p /opt/kompose
	@if ! [ -f "/opt/kompose/kompose-$$(cat ./.version_latest)$(BIN_EXT)" ]; then \
		cp "./bin/$$(cat ./.version_latest)/kompose-$(SYS_OS)-$(SYS_ARCH)$(BIN_EXT)" \
			"/opt/kompose/kompose-$$(cat ./.version_latest)$(BIN_EXT)"; \
	fi
	@if [ -f "/usr/bin/kompose$(BIN_EXT)" ]; then \
		ls -al "/usr/bin/kompose$(BIN_EXT)"; \
		printf -- "it seems like '/usr/bin/kompose$(BIN_EXT)' already exists, remove it before retrying.\n"; \
		exit 1; \
	fi
	@printf -- "requesting sudo access to create symlink at '/usr/bin/kompose$(BIN_EXT)'... "
	@sudo ln -s \
		"/opt/kompose/kompose-$$(cat ./.version_latest)$(BIN_EXT)" \
		"/usr/bin/kompose$(BIN_EXT)"

check: verify
verify:
	@which kompose \
		&& ls -al $$(which kompose) \
		&& kompose version
