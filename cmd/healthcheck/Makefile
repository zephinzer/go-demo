APP_NAME=healthcheck
PATH_TO_CERT=../../assets/certs/server.crt
PATH_TO_KEY=../../assets/keys/server.key
PORT=11114

run:
	PORT=$(PORT) \
		go run .

# run with https
runs:
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
		--file ../../deployments/Dockerfile \
		--tag zephinzer/demo-$(APP_NAME):latest \
		../../

testrun: image
	## - - - - - - - - - - - - - - - 
	## start the echoserver
	## try this out with
	## - `curl http://localhost:11113/echoserver1`
	## - `curl http://localhost:11113/echoserver2`
	## - - - - - - - - - - - - - - - 
	docker run \
		--network host \
		--env ECHOSERVER1="http://localhost:11111" \
		--env ECHOSERVER2="https://localhost:11112" \
		--env PORT=$(PORT) \
		zephinzer/demo-$(APP_NAME):latest

# testrun with https
testruns: image
	## - - - - - - - - - - - - - - - 
	## start the echoserver
	## try this out with
	## - `curl http://localhost:11113/echoserver`
	## - - - - - - - - - - - - - - - 
	docker run \
		--network host \
		--env ECHOSERVER="http://localhost:11111" \
		--env PORT=$(PORT) \
		-v "$$(pwd)/$(PATH_TO_CERT):/etc/ssl/server.crt" \
		-v "$$(pwd)/$(PATH_TO_KEY):/etc/ssl/server.key" \
		zephinzer/demo-$(APP_NAME):latest

publish: build image
	docker push zephinzer/demo-$(APP_NAME):latest
	docker tag zephinzer/demo-$(APP_NAME):latest zephinzer/demo-$(APP_NAME):$$(date +'%Y%m%d')
	docker push zephinzer/demo-$(APP_NAME):$$(date +'%Y%m%d')
	xdg-open https://hub.docker.com/r/zephinzer/demo-$(APP_NAME)
