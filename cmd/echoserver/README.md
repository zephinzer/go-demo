> [Go-Demo](../../) > [Services](../) > `echoserver`

# `echoserver`

This is a HTTP server that responds to any request URI with a standardised echo response that displays information about the request.

# Usage

Use `make run` from this directory to start the application.

The Docker image is available at https://hub.docker.com/r/zephinzer/demo-echoserver

# Config

| Environment Variable | Description |
| --- | --- |
| `HOST` | Sets the host interface to bind to (defaults to 0.0.0.0) |
| `PORT` | Sets the port to listen on (defaults to 11111) |

# More

See the [`./Makefile`](./Makefile) for other operations.
