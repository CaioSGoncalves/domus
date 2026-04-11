.DEFAULT_GOAL := build

.PHONY: fmt vet test clean generate build run

APP_NAME := domus

fmt:
	go fmt ./..
	templ fmt .

vet:
	go vet ./...

test:
	go test ./...

clean:
	rm -rf bin

generate: fmt
	templ generate

build: generate fmt vet
	mkdir -p bin
	go build -o bin/$(APP_NAME) ./cmd/$(APP_NAME)

run: generate fmt vet
	go run ./cmd/$(APP_NAME)
