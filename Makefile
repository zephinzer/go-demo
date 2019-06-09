all:
	@$(MAKE) echoserver
	@$(MAKE) fwdserver
	@$(MAKE) healthcheck

publish:
	@$(MAKE) echoserver_publish
	@$(MAKE) fwdserver_publish
	@$(MAKE) healthcheck_publish

showcase:
	@cd ./deployments/showcase \
		&& docker-compose up -V

showcase_d:
	@cd ./deployments/showcase \
		&& docker-compose up -d -V

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

########################
# provisioning recipes #
########################

dep:
	@go mod vendor

setup:
	@cd init && docker-compose up -d -V

teardown:
	@cd init && docker-compose down

###################
# utility recipes #
###################

app:
	@cd ./cmd/${APP} && $(MAKE) build && $(MAKE) image
app_publish:
	@cd ./cmd/${APP} && $(MAKE) publish

ssl:
	@mkdir -p ./assets/keys
	@openssl genrsa -out ./assets/keys/server.key 2048
	@mkdir -p ./assets/certs
	@openssl req -new -x509 -sha256 -key ./assets/keys/server.key -out ./assets/certs/server.crt -days 3650 \
		-subj "/C=SG/ST=Singapore/L=Singapore/O=Sia/CN=localhost"
