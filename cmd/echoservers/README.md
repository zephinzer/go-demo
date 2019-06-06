# `echoservers`

This is a HTTPS server that responds to any request URI with a standardised echo response that displays information about the request.

# Usage

Use `make run` from this directory to start the application.

# Config

| Environment Variable | Description |
| --- | --- |
| `HOST` | Sets the host interface to bind to (defaults to 0.0.0.0) |
| `PORT` | Sets the port to listen on (defaults to 11111) |
| `CERT` | Absolute or relative path to the server certificate file (defaults to `/etc/ssl/server.crt`) |
| `KEY` | Absolute or relative path to the server key file (defaults to `/etc/ssl/server.key`) |

# More

See the [`./Makefile`](./Makefile) for other operations.
