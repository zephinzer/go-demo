include ../../scripts/Makefile

define HELP_OUTPUT

Minikube
^^^^^^^^
Minikube is a development tool for k8s and enables you
to set up a development one-node cluster on your local
machine

Link: https://github.com/kubernetes/minikube

endef

# installs minikube locally on the machine this recipe
# is being run on
# * intended for use on a fresh cluster
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
	@curl https://api.github.com/repos/kubernetes/minikube/releases > ./.version_info

version_latest:
	@cat ./.version_info | jq '.[].tag_name' -r | sort -V | tail -n 1 > ./.version_latest
	@printf -- "latest version '$$(cat ./.version_latest)' stored in ./.version_latest\n"

download:
	@mkdir -p ./bin/$$(cat ./.version_latest)
	@curl -Lo "./bin/$$(cat ./.version_latest)/minikube-$(SYS_OS)-$(SYS_ARCH)$(BIN_EXT)" \
		"https://storage.googleapis.com/minikube/releases/$$(cat ./.version_latest)/minikube-$(SYS_OS)-$(SYS_ARCH)$(BIN_EXT)"
	@curl -Lo "./bin/$$(cat ./.version_latest)/minikube-$(SYS_OS)-$(SYS_ARCH)$(BIN_EXT).sha256" \
		"https://storage.googleapis.com/minikube/releases/$$(cat ./.version_latest)/minikube-$(SYS_OS)-$(SYS_ARCH)$(BIN_EXT).sha256"
	@chmod +x "./bin/$$(cat ./.version_latest)/minikube-$(SYS_OS)-$(SYS_ARCH)$(BIN_EXT)"

install:
	@mkdir -p /opt/minikube
	@if ! [ -f "/opt/minikube/minikube-$$(cat ./.version_latest)$(BIN_EXT)" ]; then \
		cp "./bin/$$(cat ./.version_latest)/minikube-$(SYS_OS)-$(SYS_ARCH)$(BIN_EXT)" \
			"/opt/minikube/minikube-$$(cat ./.version_latest)$(BIN_EXT)"; \
	fi
	@if [ -f "/usr/bin/minikube$(BIN_EXT)" ]; then \
		ls -al "/usr/bin/minikube$(BIN_EXT)"; \
		printf -- "it seems like '/usr/bin/minikube$(BIN_EXT)' already exists, remove it before retrying.\n"; \
		exit 1; \
	fi
	@printf -- "requesting sudo access to create symlink at '/usr/bin/minikube$(BIN_EXT)'... "
	@sudo ln -s \
		"/opt/minikube/minikube-$$(cat ./.version_latest)$(BIN_EXT)" \
		"/usr/bin/minikube$(BIN_EXT)"

verify:
	@minikube version
	@ls -al $$(which minikube)
check: verify

k8s_status:
	minikube status

MINIKUBE_MEMORY=2048
MINIKUBE_CPU=2
MINIKUBE_DISK_SIZE=8G

-include ./Makefile.properties

minikube_configure:
	#
	# set the memory capacity
	# - to modify this, create a Makefile.properties and add
	#	 	the following: MINIKUBE_MEMORY=${NEW_MEMORY_IN_MB}
	@sleep 1
	minikube config set memory $(MINIKUBE_MEMORY)
	#
	# set the cpu capacity
	# - to modify this, create a Makefile.properties and add
	#	 	the following: MINIKUBE_CPU=${NEW_CPU}
	@sleep 1
	minikube config set cpus $(MINIKUBE_CPU)
	#
	# set the disk size
	# - to modify this, create a Makefile.properties and add
	#	 	the following: MINIKUBE_DISK_SIZE=${NEW_DISK_SIZE}
	@sleep 1
	minikube config set disk-size $(MINIKUBE_DISK_SIZE)
	#
	# set the vm driver
	# - this is automatically set based on your os, to change
	#		this, simply set it yourself after this has completed
	@sleep 1
	@if [ "$(SYS_OS)" = "linux" ]; then \
		minikube config set vm-driver kvm2; \
	elif [ "$(SYS_IS)" = "darwin" ]; then \
		minikube config set vm-driver hyperkit; \
	else \
		minikube config set vm-driver none; \
	fi
	#
	# review the settings!
	@sleep 1
	@minikube config view

# this is really useful when it comes to driver issues,
# run this, and set your `vm-driver` config to another or
# try it again with different permissions
minikube_purge:
	#
	# purging the minikube installation less the binary,
	# hit ctrl + c to stop it in the next 5 seconds if you
	# ran this by accident
	#
	@sleep 1 && printf '.' && sleep 1 && printf '.' && sleep 1 && printf '.' && sleep 1 && printf '.' && sleep 1 && printf '.\nstarting...\n' && sleep 1
	@sudo rm -rf $(HOME)/.minikube

k8s_up:
	@minikube start -p minikube

k8s_down:
	@minikube stop -p minikube

k8s_rm:
	@minikube delete -p minikube
