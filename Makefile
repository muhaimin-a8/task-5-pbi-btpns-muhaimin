# postgres://{username}:{password}@{host}:{port}/postgres?sslmode=disable
psqlUri = ""

run:
	go run main.go

build:
	go build -o ./build/vix-btpns-api:1.0.0

test:
	go test ./..

migrate-up:
	migrate -database ${psqlUri} -path database/migrations up

migrate-down:
	migrate -database ${psqlUri} -path database/migrations down
