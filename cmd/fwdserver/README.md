> [Go-Demo](../../) > [Services](../) > `fwdserver`

# `fwdserver`

This is a server that forwards your requests to another server.

# Usage

Use `make run` from this directory to start the application.

The Docker image is available at https://hub.docker.com/r/zephinzer/demo-fwdserver

# Config

| Environment Variable | Description |
| --- | --- |
| `HOST` | Sets the host interface to bind to (defaults to 0.0.0.0) |
| `PORT` | Sets the port to listen on (defaults to 11111) |
| `CERT` | Absolute or relative path to the server certificate file (defaults to `/etc/ssl/server.crt`). If not found, a HTTP server will be spun up instead. |
| `KEY` | Absolute or relative path to the server key file (defaults to `/etc/ssl/server.key`). If not found, a HTTP server will be spun up instead. |

The next hop connections are configured by using environment variables corresponding to the key of the next hop server. For example, given the following when running `fwdserver`:

```sh
GOOGLE=https://google.com \
  GITHUB=https://github.com \
  PORT=1234 \
  go run .
```

Accessing Google can be done via:

```sh
curl http://localhost:1234/google
```

Accessing GitHub can be done via:

```sh
curl http://localhost:1234/github
```

# More

See the [`./Makefile`](./Makefile) for other operations.
