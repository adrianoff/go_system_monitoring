BIN := "./bin"
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

generate:
	protoc --proto_path=api --go_out=internal/pb --go-grpc_out=internal/pb api/Service.proto

build:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/app