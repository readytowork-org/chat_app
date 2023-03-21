FROM golang:1.19-alpine

# Required because go requires gcc to build
RUN apk add build-base

RUN apk add inotify-tools

RUN echo $GOPATH

RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY . /clean_web

WORKDIR /clean_web

RUN go mod tidy

ARG VERSION="4.13.0"

CMD sh /clean_web/docker/run.sh