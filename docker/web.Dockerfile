FROM golang:1.16-alpine

# Required because go requires gcc to build
RUN apk add build-base

RUN apk add inotify-tools

RUN echo $GOPATH

RUN go install github.com/go-delve/delve/cmd/dlv@latest

RUN go mod tidy

ARG VERSION="4.13.0"

# WORKDIR /clean_web

# CMD sh /clean_web/docker/run.sh