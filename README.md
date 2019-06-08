# Go Demo
A playground for Go and other things.

# Main Usage

## Setting Up

Run `make` to create all necessary binaries and Docker images.

Run `make ssl` to create the required certificates/keys to support HTTPS.

Run `make showcase` to create the Docker Compose setup.

## Try It Out

Run `curl -k https://localhost:12111/echoserver1` to get a response from `echoserver1`

Run `curl -k https://localhost:12111/echoserver2` to get a response from `echoserver2`
