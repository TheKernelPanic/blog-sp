FROM alpine:3.15.5

ARG WORK_DIRECTORY

RUN mkdir $WORK_DIRECTORY

WORKDIR $WORK_DIRECTORY

RUN apk update && \
    apk add git nano curl go glide bash

EXPOSE 80