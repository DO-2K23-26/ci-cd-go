# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.22 AS build-stage

WORKDIR /app

# COPY go.mod go.sum ./
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping


# Deploy the application binary into a lean image
FROM ubuntu:latest AS build-release-stage

WORKDIR /

COPY --from=build-stage /docker-gs-ping /docker-gs-ping

EXPOSE 8080

RUN useradd -rm -d /home/nonroot -s /bin/bash -u 1001 nonroot
USER nonroot

ENTRYPOINT ["/docker-gs-ping"]