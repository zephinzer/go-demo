# recipes here are mostly for use in the ci pipeline

include ./scripts/Makefile

define HELP_OUTPUT

Cloud Native Demos by @zephinzer
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
This repository contains tools and setups to demonstrate
cloud native technologies. The directories are structured
as such:

~/assets:      contains misc stuff like certs/keys for general use
~/deployments: example deployments for docker-compose or k8s
~/init:        setup scripts to initialise k8s clusters
~/scripts:     some tooling makefiles
~/tools:       apps for concept demonstration (written in go)

Run `make bootstrap` to setup this repository for use.

endef
export HELP_OUTPUT

# print the help message
help_root:
	@printf -- "$${HELP_OUTPUT}"

bootstrap:
	#
	# create ssl certs for tools to use https
	@sleep 3
	-@$(MAKE) ssl
	#
	# docker is required to build images
	@sleep 3
	-@$(MAKE) check_prereq_docker
	#
	# docker-compose is required to run deployments
	@sleep 3
	-@$(MAKE) check_prereq_docker_compose
	#
	# go is required to compile the tools, we check for it,
	# then go into the tools directory to retrieve dependencies
	@sleep 3
	-@$(MAKE) check_prereq_go && cd ./tools && $(MAKE) dep
	#
	# kubectl is required to communicate with the k8s cluster
	@sleep 3
	-@$(MAKE) check_prereq_kubectl
	#
	# minikube is required to bring up a k8s cluster
	@sleep 3
	-@$(MAKE) check_prereq_minikube

# deployments

build_deployments:
	# builds the images required for deployments
	@cd deployments && make build

# tools

build_tools:
	# creates binaries for all tools
	@cd tools && make all

publish_tools:
	# publishes all tools to image repository
	@cd tools && make publish
