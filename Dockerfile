# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.22 AS build-stage

WORKDIR /app

# COPY go.mod go.sum ./
COPY go.mod ./
COPY go.sum ./

RUN go mod download

# Copy .go files, modules
COPY main.go ./

COPY models/*.go ./models/
COPY routes/*.go ./routes/
COPY services/*.go ./services/
COPY controllers/*.go ./controllers/
COPY database/*.go ./database/

RUN CGO_ENABLED=0 GOOS=linux go build -o /goserver


# Deploy the application binary into a lean image
FROM ubuntu:latest AS build-release-stage

WORKDIR /

COPY --from=build-stage /goserver /goserver
COPY cities.json ./

EXPOSE 2022

RUN useradd -rm -d /home/nonroot -s /bin/bash -u 1001 nonroot
USER nonroot

ENTRYPOINT ["/goserver"]