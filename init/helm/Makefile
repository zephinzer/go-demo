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
	@$(MAKE) install_helm
	@$(MAKE) install_tiller

install_helm:
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

install_tiller:
	@mkdir -p /opt/tiller
	@if ! [ -f "/opt/tiller/tiller-$$(cat ./.version_latest)$(BIN_EXT)" ]; then \
		cp "./bin/$$(cat ./.version_latest)/tiller-$(SYS_OS)-$(SYS_ARCH)$(BIN_EXT)" \
			"/opt/tiller/tiller-$$(cat ./.version_latest)$(BIN_EXT)"; \
	fi
	@if [ -f "/usr/bin/tiller$(BIN_EXT)" ]; then \
		ls -al "/usr/bin/tiller$(BIN_EXT)"; \
		printf -- "it seems like '/usr/bin/tiller$(BIN_EXT)' already exists, remove it before retrying.\n"; \
		exit 1; \
	fi
	@printf -- "requesting sudo access to create symlink at '/usr/bin/tiller$(BIN_EXT)'... "
	@sudo ln -s \
		"/opt/tiller/tiller-$$(cat ./.version_latest)$(BIN_EXT)" \
		"/usr/bin/tiller$(BIN_EXT)"

check: verify
verify:
	-@helm version && ls -al $$(which helm)
	-@tiller version && ls -al $$(which tiller)

# note: for production, use a propert root ca
helm_create_cert_authority:
	@mkdir -p ./ssl
	@openssl genrsa \
		-out ./ssl/ca.key.pem 4096
	@openssl req -new -x509 -sha256 \
		-key ./ssl/ca.key.pem \
		-days 3650 \
		-extensions v3_ca \
		-subj "/C=SG/ST=Singapore/L=Singapore/O=zephinzer-demo/OU=helmca/CN=helmca" \
		-out ./ssl/ca.cert.pem

# note: for production, create a cert for each user of helm
helm_create_user_keys:
	@mkdir -p ./ssl
	@openssl genrsa \
		-out ./ssl/helm.key.pem 4096
	@openssl genrsa \
		-out ./ssl/tiller.key.pem 4096

helm_create_user_csrs:
	@openssl req -new -sha256 \
		-key ./ssl/helm.key.pem \
		-subj "/C=SG/ST=Singapore/L=Singapore/O=zephinzer-demo/OU=helm/CN=helmuser" \
		-out ./ssl/helm.csr.pem
	@openssl req -new -sha256 \
		-key ./ssl/tiller.key.pem \
		-subj "/C=SG/ST=Singapore/L=Singapore/O=zephinzer-demo/OU=tiller/CN=tilleruser" \
		-out ./ssl/tiller.csr.pem

helm_create_user_certs:
	@openssl x509 -req \
		-days 3650 \
		-CAcreateserial \
		-CA ./ssl/ca.cert.pem \
		-CAkey ./ssl/ca.key.pem \
		-in ./ssl/helm.csr.pem \
		-out ./ssl/helm.cert.pem
	@openssl x509 -req \
		-days 3650 \
		-CAcreateserial \
		-CA ./ssl/ca.cert.pem \
		-CAkey ./ssl/ca.key.pem \
		-in ./ssl/tiller.csr.pem \
		-out ./ssl/tiller.cert.pem

helm_init:
	@$(MAKE) _helm_init FLAGS="--dry-run --debug"
	@printf -- "use 'make helm_init_do' to apply it (have you checked out 'make helm_secure_init'?)\n"
helm_init_do:
	@$(MAKE) _helm_init
_helm_init:
	@helm init ${FLAGS} --history-max 500

helm_secure_init:
	@$(MAKE) _helm_secure_init FLAGS="--dry-run --debug"
	@printf -- "use 'make helm_secure_init_do' to apply it\n"
helm_secure_init_do:
	@$(MAKE) _helm_secure_init
_helm_secure_init:
	@if kubectl get sa --all-namespaces | grep tiller; then \
		helm init ${FLAGS} \
			--history-max 500 \
			--service-account=tiller \
			--tiller-namespace=tiller \
			--tiller-tls \
			--tiller-tls-cert ./ssl/tiller.cert.pem \
			--tiller-tls-key ./ssl/tiller.key.pem \
			--tiller-tls-verify \
			--tls-ca-cert ./ssl/ca.cert.pem; \
	else \
		printf -- "service account 'tiller' is not available - run 'kubectl apply -f ./resources/helm_namespace_access.yaml' first\n"; \
	fi