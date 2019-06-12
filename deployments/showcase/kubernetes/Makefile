all:
	@minikube status
	@$(MAKE) _namespace
	@$(MAKE) echoserver
	@$(MAKE) fwdserver
	@$(MAKE) healthcheck
	@$(MAKE) traffic_generator

_namespace:
	@kubectl apply -f ./namespace.yaml

echoserver:
	@kubectl apply -f ./echoserver.yaml
	@kubectl apply -f ./echoserver1.yaml
	@kubectl apply -f ./echoserver2.yaml
	@kubectl apply -f ./echoserver3.yaml

fwdserver:
	@kubectl apply -f ./fwdserver.yaml

healthcheck:
	@kubectl apply -f ./healthcheck.yaml

traffic_generator:
	@kubectl apply -f ./traffic-generator.yaml
