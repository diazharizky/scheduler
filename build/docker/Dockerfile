FROM golang:1.13-alpine as builder

ARG git_tag
ARG git_commit

RUN apk add --no-cache git build-base curl
RUN go get -u github.com/gobuffalo/packr/v2/packr2

ADD . /go/src/github.com/diazharizky/scheduler 
WORKDIR /go/src/github.com/diazharizky/scheduler
RUN packr2
RUN go mod vendor
RUN go build github.com/diazharizky/scheduler
RUN packr2 clean

FROM alpine

COPY --from=builder /go/src/github.com/diazharizky/scheduler/scheduler /usr/bin/scheduler

ENTRYPOINT ["scheduler"]
