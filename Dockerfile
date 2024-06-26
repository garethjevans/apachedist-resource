FROM golang:1.22.4 as go
RUN GOPROXY=direct GO111MODULES=on go install github.com/garethjevans/apachedist-resource@main

FROM ubuntu:24.04
RUN apt-get update && apt-get install -y ca-certificates && apt-get clean

COPY --from=go /go/bin/apachedist-resource /bin/apachedist-resource

RUN apachedist-resource --help

COPY scripts/in /opt/resource/in
COPY scripts/check /opt/resource/check
