FROM golang:alpine

ENV GO111MODULE=on
ENV CGO_ENABLED=0

RUN go get github.com/stretchr/testify/assert
