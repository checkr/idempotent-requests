.PHONY: vendor fmt tests

vendor:
	go mod tidy
	go mod vendor

fmt:
	go fmt ./...

test:
	go test ./pkg/...