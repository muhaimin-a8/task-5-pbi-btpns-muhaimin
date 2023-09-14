FROM golang:1.20-alpine AS builder
WORKDIR app

COPY . .

RUN go mod download
RUN go build -o ./pbi-btpns-api:v1.0.0

FROM golang:1.20-alpine as publisher
WORKDIR app
COPY --from=builder
