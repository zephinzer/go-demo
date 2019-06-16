include ../scripts/Makefile

define HELP_OUTPUT

Example deployments
^^^^^^^^^^^^^^^^^^^
This directory contains example deployment manifests for
Docker Compose and Kubernetes.

endef
export HELP_OUTPUT

# print the help message
help_deployments:
	@printf -- "$${HELP_OUTPUT}"

prereq_check:
	# checks if pre-requsites have been met
	-@$(MAKE) check_prereq_docker
	-@$(MAKE) check_prereq_docker_compose
	-@$(MAKE) check_prereq_kubectl

build:
	# builds images for all deployments
	@$(MAKE) showcase_build

showcase:
	# runs the showcase deployment in the foreground
	@$(MAKE) _showcase ARGS="up -V"
showcase_d:
	# runs the showcase deployment in the background
	@$(MAKE) _showcase ARGS="up -d -V"
showcase_build:
	# builds the images for the showcase deployment
	@cd ../tools && $(MAKE) all
	@$(MAKE) _showcase ARGS="build"
_showcase:
	@cd ./showcase/docker-compose && docker-compose ${ARGS}