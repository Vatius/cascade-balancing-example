BINARY_PATH = ./bin

.PHONY: build
build:
	mkdir -p ./bin
	go build -v -o $(BINARY_PATH)/server ./cmd/server

.PHONY: test
test:
	go test

.DEFAULT_GOAL := build