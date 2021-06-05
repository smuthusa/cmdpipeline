.DEFAULT_GOAL := default
default: build

BINARY_NAME = pipeline

build:
	go build -o $(BINARY_NAME) -v $(BINARY_NAME).go

linux-build:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME) -v $(BINARY_NAME).go

run: build
	go run -v $(BINARY_NAME).go execute --file ./resources/tasks.yaml --validate