all:
	@$(MAKE) echoserver
	@$(MAKE) fwdserver

publish:
	@$(MAKE) echoserver_publish
	@$(MAKE) fwdserver_publish

showcase:
	@cd ./deployments/showcase \
		&& docker-compose up -V

echoserver:
	@$(MAKE) app APP=echoserver
echoserver_publish: echoserver
	@$(MAKE) app_publish APP=echoserver

fwdserver:
	@$(MAKE) app APP=fwdserver
fwdserver_publish: fwdserver
	@$(MAKE) app_publish APP=fwdserver

########################
# provisioning recipes #
########################

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
