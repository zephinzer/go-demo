FROM alpine:3.9.4 AS dev
ARG BIN_NAME="app"
WORKDIR /
RUN apk update --no-cache \
  && apk upgrade --no-cache \
  && apk add ca-certificates --no-cache \
  && update-ca-certificates \
  && adduser -D -u 500 -H app
COPY ./build/bin/${BIN_NAME} /app
RUN chown app:app /app
USER app
ENTRYPOINT [ "/app" ]

FROM scratch AS production
COPY --from=dev /etc/ssl/certs/ /etc/ssl/certs/
COPY --from=dev /etc/passwd /etc/passwd
COPY --from=dev /app /app
USER app
ENTRYPOINT [ "/app" ]
ARG BIN_NAME="app"
ARG DESCRIPTION="[not set] a demo app written in go for code/service showcase purposes"
LABEL maintainer="@zephinzer" \
  description="${DESCRIPTION}" \
  maintainer_url="https://github.com/zephinzer" \
  repository_url="https://github.com/zephinzer/go-demo" \
  container_url="https://hub.docker.com/r/zephinzer/demo-${BIN_NAME}" \
  application_url="https://github.com/zephinzer/go-demo/tree/master/cmd/${BIN_NAME}"
