BINARY_PATH = ./bin

.PHONY: build
build:
	mkdir -p ./bin
	go build -v -o $(BINARY_PATH)/server ./cmd/server
	go build -v -o $(BINARY_PATH)/client ./cmd/client

.PHONY: test
test:
	go test -v ./cmd/server

.DEFAULT_GOAL := build