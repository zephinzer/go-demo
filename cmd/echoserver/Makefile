APP_NAME=echoserver
PORT=11111

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
	## try this out with `curl http://localhost:11111`
	## - - - - - - - - - - - - - - - 
	docker run \
		--network host \
		--env PORT=$(PORT) \
		zephinzer/demo-$(APP_NAME):latest

publish: build image
	docker push zephinzer/demo-$(APP_NAME):latest
	xdg-open https://hub.docker.com/r/zephinzer/demo-$(APP_NAME)

apply:
	if minikube status ; then kubectl apply -f ./k8s.yaml; else echo 'no minikube instance was found running'; fi
