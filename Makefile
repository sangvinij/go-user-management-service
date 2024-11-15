build:
	go build -v ./src/cmd/

run:
	go run ./src/cmd

lint:
	golangci-lint run

format:
	go fmt ./...
	swag fmt .

docker-build:
	docker compose build

docker-run:
	docker compose up
