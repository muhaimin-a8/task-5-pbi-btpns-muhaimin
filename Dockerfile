FROM golang:1.20-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./build/pbi-btpns-api:v1.0.0

# Production stage
FROM scratch

ENV STAGE=production

COPY --from=builder /app/build/pbi-btpns-api:v1.0.0 .

EXPOSE 8080
ENTRYPOINT ["./pbi-btpns-api:v1.0.0"]
