# Base Image for building
# build is run with `docker-compose up generator && docker-compose up builder` for better caching with local gopath
FROM rucciva/golang:1.15.2-alpine as base

# Final Image
# run `docker-compose up generator && docker-compose up builder` first
FROM alpine:3.10.2

RUN apk add --no-cache ca-certificates

COPY bin/scheduler /usr/local/bin

ENTRYPOINT /usr/local/bin/scheduler
