GOPATH:=$(shell go env GOPATH)

.PHONY: init
# init env
init:
	go install github.com/google/wire/cmd/wire@latest


.PHONY: build
# build
build:
	go mod tidy&&mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: generate
# generate
generate:
	go generate ./...
