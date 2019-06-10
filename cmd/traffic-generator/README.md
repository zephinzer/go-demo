> [Go-Demo](../../) > [Services](../) > `traffic-generator`


# `traffic-generator`


This service is a load generator to send mock loads to a target server of choice.


# Usage


Use `make run` from this directory to start the application.

The Docker image is available at https://hub.docker.com/r/zephinzer/demo-traffic-generator


# Config


| Environment Variable | Description |
| --- | --- |
| `BODY` | The body data for the requester to use to send to the remote server specified at `URL` |
| `METHOD` | The HTTP method which will be used to send a request to the remote server |
| `REQ_CONCURRECY` | The number of concurrent requests being made |
| `REQ_INTERVAL` | The duration between which requests should be made |
| `REQ_RATE` | The rate of requests made by each requester thread |
| `URL` | The URL to send requests to |


# More


See the [`./Makefile`](./Makefile) for other operations.
