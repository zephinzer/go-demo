APP_NAME=echoserver

run:
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
	## try this out with `curl http://localhost:11111`
	## - - - - - - - - - - - - - - - 
	docker run \
		-p 11111:11111 \
		zephinzer/demo-$(APP_NAME):latest

publish: build image
	docker push zephinzer/demo-$(APP_NAME):latest
