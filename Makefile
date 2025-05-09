all: test

vendor:
	go mod vendor

test:
	go test -coverprofile=coverage.out ./...

lint:
	golangci-lint run ./...

.PHONY: test
