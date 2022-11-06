# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

ARG PROJECT_BINARY=bestir-dev-client
ARG PROJECT_BUILD_DIR=./build/bin

WORKDIR /app

RUN ls

COPY go.mod ./
COPY go.sum ./
RUN go mod download

# temporary measure
COPY *.go ./

RUN go build -o /bestir-dev-client

EXPOSE 8080

 CMD [ "/bestir-dev-client" ]