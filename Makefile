.PHONY: build
build:
	go build -v ./src/cmd/

.PHONY: run
run:
	go run ./src/cmd

.PHONY: docker-build
docker-run:
	docker compose build

.PHONY: docker-run
docker-run:
	docker compose up

.DEFAULT_GOAL := run
