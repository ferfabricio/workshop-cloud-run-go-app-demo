# syntax=docker/dockerfile:1
##
## BUILD
##

FROM golang:1.23 AS base

ENV GOOS linux
ENV CGO_ENABLED 0

WORKDIR /workspace

COPY . .

RUN go mod download
RUN go build -o http-app ./cmd/main.go

##
## DEPLOY
##

FROM alpine:latest

WORKDIR /

COPY --from=base /workspace/http-app ./app

ENTRYPOINT ["/app"]