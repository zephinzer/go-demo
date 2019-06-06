all:
	@$(MAKE) echoserver
	@$(MAKE) echoservers

echoserver:
	@cd ./cmd/echoserver && $(MAKE) build && $(MAKE) image
echoserver_publish: echoserver
	@cd ./cmd/echoserver && $(MAKE) publish

echoservers:
	@cd ./cmd/echoservers && $(MAKE) build && $(MAKE) image
echoservers_publish: echoservers
	@cd ./cmd/echoservers && $(MAKE) publish

init:
	@cd init && docker-compose up -d -V

denit:
	@cd init && docker-compose down

ssl:
	@mkdir -p ./assets/keys
	@openssl genrsa -out ./assets/keys/server.key 2048
	@mkdir -p ./assets/certs
	@openssl req -new -x509 -sha256 -key ./assets/keys/server.key -out ./assets/certs/server.crt -days 3650 \
		-subj "/C=SG/ST=Singapore/L=Singapore/O=Sia/CN=localhost"
