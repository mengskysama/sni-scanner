FROM golang:1.14.3-buster as builder

COPY go.mod /build/
COPY scanner /build/scanner
COPY main.go /build/
RUN cd /build/ && go test
RUN cd /build/ && go build
