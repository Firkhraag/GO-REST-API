FROM golang:1.14 as build-env

WORKDIR /go/src/github.com/firkhraag/apiserver
ADD . /go/src/github.com/firkhraag/apiserver

ENV GO111MODULE=on

RUN go mod download
RUN go mod verify
RUN go install ./cmd/apiserver
