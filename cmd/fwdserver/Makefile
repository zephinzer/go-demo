APP_NAME=fwdserver
PORT=11113

run:
	PORT=$(PORT) go run .

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

publish: build image
	docker push zephinzer/demo-$(APP_NAME):latest
