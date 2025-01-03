BINARY_NAME=ethparser

.PHONY: build

PACKAGE=github.com/guneyin/ethparser

init: clean mod tidy vet build

clean:
	go clean
	rm -f ${BINARY_NAME}

mod:
	go mod download

tidy:
	go mod tidy

vet:
	go vet ./...

build:
	go build -o ${BINARY_NAME} .

test:
	go test ./...

cover:
	go test ./... -coverprofile=cover.out
	go tool cover -html=cover.out

run:
	go run .
