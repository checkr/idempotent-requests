.PHONY: build vendor fmt tests

build:
	@mkdir -p ./bin
	go build -o ./bin/idempotent-requests-server

vendor:
	go mod tidy
	go mod vendor

fmt:
	go fmt ./...

test:
	go test -race -covermode=atomic ./pkg/...