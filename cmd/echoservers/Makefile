APP_NAME=echoservers
PATH_TO_CERT=../../assets/certs/server.crt
PATH_TO_KEY=../../assets/keys/server.key

run:
	CERT=$(PATH_TO_CERT) \
		KEY=$(PATH_TO_KEY) \
		go run .

build:
	CGO_ENABLED=0 go build -a -ldflags "-extldflags '-static'" -o ../../build/bin/$(APP_NAME) .

image:
	docker build \
		--build-arg BIN_NAME=$(APP_NAME) \
		--file ../../deployments/Dockerfile \
		--tag zephinzer/demo-$(APP_NAME):latest \
		../../build/bin

testrun: image
	## - - - - - - - - - - - - - - - 
	## try this out with `curl -k https://localhost:11111`
	## - - - - - - - - - - - - - - - 
	docker run \
		-v "$$(pwd)/$(PATH_TO_CERT):/etc/ssl/server.crt" \
		-v "$$(pwd)/$(PATH_TO_KEY):/etc/ssl/server.key" \
		-p 11111:11111 \
		zephinzer/demo-$(APP_NAME):latest

publish: build image
	docker push zephinzer/demo-$(APP_NAME):latest
