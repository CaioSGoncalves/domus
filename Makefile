.DEFAULT_GOAL := build


APP_NAME := domus

.PHONY: fmt vet test
fmt:
	go fmt ./..
	templ fmt .

vet:
	go vet ./...

test:
	go test ./...

.PHONY: clean run build-front build build-linux
clean:
	rm -rf bin

run: fmt vet
	go run ./cmd/$(APP_NAME)

build-front:
	cd ./frontend && npm run build

build: fmt vet build-front
	mkdir -p bin
	go build -o bin/$(APP_NAME) ./cmd/$(APP_NAME)

build-linux: fmt vet build-front
	mkdir -p bin
	GOOS=linux GOARCH=amd64 go build -o bin/$(APP_NAME) ./cmd/$(APP_NAME)
