install: # make sure the purge recipe is the reverse of this
	@minikube status -v8 -p ${NAME}
	@$(MAKE) _namespace
	@$(MAKE) echoserver
	@$(MAKE) fwdserver
	@$(MAKE) healthcheck
	@$(MAKE) traffic_generator

purge: # make sure this is the reverse order from above
	@minikube status -v8 -p ${NAME}
	@$(MAKE) purge_traffic_generator
	@$(MAKE) purge_healthcheck
	@$(MAKE) purge_fwdserver
	@$(MAKE) purge_echoserver
	@$(MAKE) _purge_namespace

_namespace:
	@kubectl apply -f ./namespace.yaml
_purge_namespace:
	@kubectl delete -f ./namespace.yaml

echoserver:
	@kubectl apply -f ./echoserver.yaml
	@kubectl apply -f ./echoserver1.yaml
	@kubectl apply -f ./echoserver2.yaml
	@kubectl apply -f ./echoserver3.yaml
purge_echoserver:
	@kubectl delete -f ./echoserver3.yaml
	@kubectl delete -f ./echoserver2.yaml
	@kubectl delete -f ./echoserver1.yaml
	@kubectl delete -f ./echoserver.yaml

fwdserver:
	@kubectl apply -f ./fwdserver.yaml
purge_fwdserver:
	@kubectl delete -f ./fwdserver.yaml

healthcheck:
	@kubectl apply -f ./healthcheck.yaml
purge_healthcheck:
	@kubectl delete -f ./healthcheck.yaml

traffic_generator:
	@kubectl apply -f ./traffic-generator.yaml
	@kubectl apply -f ./traffic-generator1.yaml
	@kubectl apply -f ./traffic-generator2.yaml
	@kubectl apply -f ./traffic-generator3.yaml
purge_traffic_generator:
	@kubectl delete -f ./traffic-generator3.yaml
	@kubectl delete -f ./traffic-generator2.yaml
	@kubectl delete -f ./traffic-generator1.yaml
	@kubectl delete -f ./traffic-generator.yaml