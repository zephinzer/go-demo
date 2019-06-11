> [Go-Demo](../../../) > [Deployments](../../) > [Showcase](../) > Docker Compose

# Showcase (Docker Compose)

This directory contains a Docker Compose that brings up a set of services where each service in `go-demo` is represented:

| Service | Accessible At |
| --- | --- |
| `echoserver1` | http://localhost:11111 |
| `echoserver2` | http://localhost:11112 |
| `echoserver3` | http://localhost:11113 |
| `fwdserver` | http://localhost:11211 |
| `healthcheck` | http://localhost:11311 |
| `traffic-generator` | http://localhost:11411 |

To start it, use:

```sh
docker-compose up -V
```
