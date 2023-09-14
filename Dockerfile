FROM golang:1.20-alpine AS builder
WORKDIR app

COPY . .

RUN go mod download
RUN go build -o ./build/pbi-btpns-api:v1.0.0

FROM scratch

ENV STAGE=production

WORKDIR /
COPY --from=builder ./build/pbi-btpns-api:v1.0.0 ./pbi-btpns-api:v1.0.0

EXPOSE 8080
ENTRYPOINT ["./pbi-btpns-api:v1.0.0"]
