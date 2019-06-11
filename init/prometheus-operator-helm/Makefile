include ../Makefile

prerequisites:
	which helm && helm version --client

install:
	helm install stable/prometheus-operator

download:
	helm get stable/prometheus-operator
