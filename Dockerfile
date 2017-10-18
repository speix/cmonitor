FROM golang

ADD . /go/src/github.com/speix/cmonitor

WORKDIR /go/src/github.com/speix/cmonitor

ARG CM_API_KEY
ARG CM_API_CLIENT

ENV CM_API_KEY $CM_API_KEY
ENV CM_API_CLIENT $CM_API_CLIENT

RUN go install github.com/speix/cmonitor

ENTRYPOINT /go/bin/cmonitor

EXPOSE 8080