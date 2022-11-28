GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

.PHONY: init
# init env
init:
	go install github.com/google/wire/cmd/wire@latest


.PHONY: generate
# generate
generate:
	go generate ./...

.PHONY: build
# build
build:
	go mod tidy&&mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: docker
# build
docker:
	go mod tidy&&mkdir -p bin/ && CGO_ENABLED=0 GOOS=linux go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./... \
	&& docker build . -t xuxiaoyu/blog2:$(VERSION) \
	&& docker push xuxiaoyu/blog2:$(VERSION);
