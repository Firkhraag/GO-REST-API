.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: format
format:
	go fmt ./...

.PHONY: migrate
migrate:
	go run cmd/migrations/migrations.go $(dbname)

.DEFAULT_GOAL := build
