include ../scripts/Makefile

define HELP_OUTPUT

The '~/tools' directory contains Go applications used for
demonstrating acitivites/concepts in cloud-native setups.

The following is a summary of the custom services:
- echoserver
- fwdserver
- healthcheck
- traffic-generator

Run any of the above as make recipes to build the binary and
image (eg. `make echoserver` to build the echoserver service).

Alternatively to build everything: `make all`

endef
export HELP_OUTPUT

help_tools:
	@echo "$${HELP_OUTPUT}"

all:
	@$(MAKE) echoserver
	@$(MAKE) fwdserver
	@$(MAKE) healthcheck
	@$(MAKE) traffic-generator

publish:
	@$(MAKE) echoserver_publish
	@$(MAKE) fwdserver_publish
	@$(MAKE) healthcheck_publish
	@$(MAKE) traffic-generator_publish

echoserver: dep
	@$(MAKE) app APP=echoserver
echoserver_publish: echoserver
	@$(MAKE) app_publish APP=echoserver

fwdserver: dep
	@$(MAKE) app APP=fwdserver
fwdserver_publish: fwdserver
	@$(MAKE) app_publish APP=fwdserver

healthcheck: dep
	@$(MAKE) app APP=healthcheck
healthcheck_publish: healthcheck
	@$(MAKE) app_publish APP=healthcheck

traffic-generator: dep
	@$(MAKE) app APP=traffic-generator
traffic-generator_publish: traffic-generator
	@$(MAKE) app_publish APP=traffic-generator

###################
# utility recipes #
###################

app:
	# builds the tools
	@cd ./cmd/${APP} && $(MAKE) build && $(MAKE) image
app_publish:
	# publishes the tools
	@cd ./cmd/${APP} && $(MAKE) publish

dep:
	# installs the go dependencies
	@GO111MODULE=on go mod vendor -v
