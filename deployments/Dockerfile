FROM alpine:3.9.4 AS dev
ARG BIN_NAME="app"
WORKDIR /
RUN apk update --no-cache \
  && apk upgrade --no-cache \
  && apk add ca-certificates --no-cache \
  && update-ca-certificates \
  && adduser -D -u 500 -H app
COPY ./${BIN_NAME} /app
RUN chown app:app /app
USER app
ENTRYPOINT [ "/app" ]

FROM scratch AS production
COPY --from=dev /etc/ssl/certs/ /etc/ssl/certs/
COPY --from=dev /etc/passwd /etc/passwd
COPY --from=dev /app /app
USER app
ENTRYPOINT [ "/app" ]