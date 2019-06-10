include ../Makefile

version_info:
	@curl https://api.github.com/repos/helm/helm/releases > ./.version_info

version_latest:
	@cat ./.version_info | jq '.[].tag_name' -r | egrep "^v[0-9]+\.[0-9]+\.[0-9]+$$" | sort -V | tail -n 1 > ./.version_latest
	@printf -- "latest version '$$(cat ./.version_latest)' stored in ./.version_latest\n"

download:
	@mkdir -p ./bin/$$(cat ./.version_latest)
	@curl -Lo "./bin/$$(cat ./.version_latest)/helm-$(SYS_OS)-$(SYS_ARCH).tar.gz" \
		"https://get.helm.sh/helm-$$(cat ./.version_latest)-$(SYS_OS)-$(SYS_ARCH).tar.gz"
	@curl -Lo "./bin/$$(cat ./.version_latest)/helm-$(SYS_OS)-$(SYS_ARCH).tar.gz.sha256" \
		"https://get.helm.sh/helm-$$(cat ./.version_latest)-$(SYS_OS)-$(SYS_ARCH).tar.gz.sha256"
	@cd ./bin/$$(cat ./.version_latest) \
		&& tar -xzvf "./helm-$(SYS_OS)-$(SYS_ARCH).tar.gz" \
		&& mv ./$(SYS_OS)-$(SYS_ARCH)/* . \
		&& rm -rf ./$(SYS_OS)-$(SYS_ARCH) \
		&& mv ./helm$(BIN_EXT) "./helm-$(SYS_OS)-$(SYS_ARCH)$(BIN_EXT)" \
		&& mv ./tiller$(BIN_EXT) "./tiller-$(SYS_OS)-$(SYS_ARCH)$(BIN_EXT)"

install:
	@mkdir -p /opt/helm
	@if ! [ -f "/opt/helm/helm-$$(cat ./.version_latest)$(BIN_EXT)" ]; then \
		cp "./bin/$$(cat ./.version_latest)/helm-$(SYS_OS)-$(SYS_ARCH)$(BIN_EXT)" \
			"/opt/helm/helm-$$(cat ./.version_latest)$(BIN_EXT)"; \
	fi
	@if [ -f "/usr/bin/helm$(BIN_EXT)" ]; then \
		ls -al "/usr/bin/helm$(BIN_EXT)"; \
		printf -- "it seems like '/usr/bin/helm$(BIN_EXT)' already exists, remove it before retrying.\n"; \
		exit 1; \
	fi
	@printf -- "requesting sudo access to create symlink at '/usr/bin/helm$(BIN_EXT)'... "
	@sudo ln -s \
		"/opt/helm/helm-$$(cat ./.version_latest)$(BIN_EXT)" \
		"/usr/bin/helm$(BIN_EXT)"

verify:
	-@helm version
	@ls -al $$(which helm)
check: verify

helm_init:
	helm init --history-max 200
