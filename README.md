# Go Demo
A playground for Go and other things like container orchestration.

# Main Usage

## Setting Up

Run `make` to create all necessary binaries and Docker images.

Run `make ssl` to create the required certificates/keys to support HTTPS.

Run `make showcase` to create the Docker Compose setup.

## Try It Out

### `echoserver`

Run `curl -k https://localhost:11111` to get a response from `echoserver1`

Run `curl -k https://localhost:11112` to get a response from `echoserver1`

### `fwdserver`

Run `curl -k https://localhost:12111/echoserver1` to get a proxied response from `echoserver1`

Run `curl -k https://localhost:12111/echoserver2` to get a proxied response from `echoserver2`

### `healthcheck`

Run `curl -k https://localhost:11114/liveness` to get a liveness check response from `healthcheck`

Run `curl -k https://localhost:11114/liveness/false` to set the liveness check response from `healthcheck` to failing

Run `curl -k https://localhost:11114/readiness` to get a readiness check response from `healthcheck`

Run `curl -k https://localhost:11114/readiness/false` to set the readiness check response from `healthcheck` to failing

# License
Feel free to use these examples in your own workshops/tutorials!

Code and usage of its resulting binaries is licensed under the permissive [MIT license](./LICENSE)

Content is licensed under the [Creative Commons Attribution-ShareAlike 3.0 (CC BY-SA 3.0 SG) license](https://creativecommons.org/licenses/by-sa/3.0/sg/). For attribution, use the following:

## HTML

```html
Content was created by <a href="https://github.com/zephinzer" target="_blank">Joseph Matthias Goh/@zephinzer</a> and the original content can be found at <a href="https://github.com/zephinzer/go-demo" target="_blank">https://github.com/zephinzer/go-demo</a>.
```

## Markdown

```markdown
Content was created by [Joseph Matthias Goh/@zephinzer](https://github.com/zephinzer) and the original content can be found at [https://github.com/zephinzer/go-demo](https://github.com/zephinzer/go-demo).
```

# Cheers!
