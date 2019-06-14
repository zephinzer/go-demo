> [Go-Demo](../../) > [Services](../) > `healthcheck`


# `healthcheck`


This service exposes 2 endpoints meant for demonstrating the effects of liveness/readiness checks on a container orchestrator. This service is intended for deployment onto a Kubernetes cluster for demonstration of pods destruction when not alive (liveness check failing), and re-routing of network requests away from pods delcaring themselves not ready (readiness check failing).


# Usage


Use `make run` from this directory to start the application.

The Docker image is available at https://hub.docker.com/r/zephinzer/demo-healthcheck

The GitHub repository is available at https://github.com/zephinzer/go-demo

To deploy this onto a Kubernetes cluster, run `kubectl apply -f ./k8s.yaml`. To modify the envrionment variable, change the ConfigMap segment in the manifest and re-run the `kubectl apply`.


# Config


| Environment Variable | Description |
| --- | --- |
| `HOST` | Sets the host interface to bind to (defaults to 0.0.0.0) |
| `PORT` | Sets the port to listen on (defaults to 11111) |
| `CERT` | Absolute or relative path to the server certificate file (defaults to `/etc/ssl/server.crt`). If not found, a HTTP server will be spun up instead. |
| `KEY` | Absolute or relative path to the server key file (defaults to `/etc/ssl/server.key`). If not found, a HTTP server will be spun up instead. |
| `ALIVE` | Sets the initial liveness status to alive or otherwise. Defaults to `true` on startup |
| `READY` | Sets the initial readiness status to ready or otherwise. Defaults to `true` on startup |


# Endpoints


| Endpoint | Description |
| --- | --- |
| `/liveness` | Retrieves the liveness status. Returns 200 on okay, 500 on errors |
| `/liveness/{status}` | Sets the liveness status. `status` can be `1`/`0` or `true`/`false`. Returns 200 on okay, 500 on errors |
| `/readiness` | Retrieves the liveness status. Returns 200 on okay, 500 on errors |
| `/readiness/{status}` | Sets the liveness status. `status` can be `1`/`0` or `true`/`false`. Returns 200 on okay, 500 on errors |


# More


See the [`./Makefile`](./Makefile) for other operations.
