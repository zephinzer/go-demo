APP_NAME=traffic-generator

# print the help message
help:
	@less ./README.md

run:
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
	docker run \
		--network host \
		zephinzer/demo-$(APP_NAME):latest

publish: build image
	docker push zephinzer/demo-$(APP_NAME):latest
	docker tag zephinzer/demo-$(APP_NAME):latest zephinzer/demo-$(APP_NAME):$$(date +'%Y%m%d')
	docker push zephinzer/demo-$(APP_NAME):$$(date +'%Y%m%d')
	xdg-open https://hub.docker.com/r/zephinzer/demo-$(APP_NAME)
