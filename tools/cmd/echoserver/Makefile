APP_NAME=echoserver
PATH_TO_CERT=../../../assets/certs/server.crt
PATH_TO_KEY=../../../assets/keys/server.key
PORT=11111

run:
	CERT=$(PATH_TO_CERT) \
		KEY=$(PATH_TO_KEY) \
		PORT=$(PORT) \
		go run .

build:
	CGO_ENABLED=0 go build -a -ldflags "-extldflags '-static'" -o ../../build/bin/$(APP_NAME) .

image:
	docker build \
		--build-arg DESCRIPTION="$$(cat ./README.md)" \
		--build-arg BIN_NAME=$(APP_NAME) \
		--file ../../Dockerfile \
		--tag zephinzer/demo-$(APP_NAME):latest \
		../../

testrun: image
	## - - - - - - - - - - - - - - - 
	## try this out with `curl http://localhost:11111`
	## - - - - - - - - - - - - - - - 
	docker run \
		--network host \
		--env PORT=$(PORT) \
		zephinzer/demo-$(APP_NAME):latest

testruns: image
	## - - - - - - - - - - - - - - - 
	## try this out with `curl -k https://localhost:11111`
	## - - - - - - - - - - - - - - - 
	docker run \
		--network host \
		--env PORT=$(PORT) \
		-v "$$(pwd)/$(PATH_TO_CERT):/etc/ssl/server.crt" \
		-v "$$(pwd)/$(PATH_TO_KEY):/etc/ssl/server.key" \
		zephinzer/demo-$(APP_NAME):latest

publish: build image
	docker push zephinzer/demo-$(APP_NAME):latest
	docker tag zephinzer/demo-$(APP_NAME):latest zephinzer/demo-$(APP_NAME):$$(date +'%Y%m%d')
	docker push zephinzer/demo-$(APP_NAME):$$(date +'%Y%m%d')
	xdg-open https://hub.docker.com/r/zephinzer/demo-$(APP_NAME)
