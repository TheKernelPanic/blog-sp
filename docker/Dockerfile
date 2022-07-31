FROM alpine:3.15.5

ARG WORK_DIRECTORY

ENV GOPATH=/usr/lib/go/bin
ENV GOROOT=/usr/lib/go/

RUN mkdir $WORK_DIRECTORY

WORKDIR $WORK_DIRECTORY

RUN apk update && \
    apk add git nano curl go glide bash

EXPOSE 80