BIN_MONITORING := "./bin/monitoring"
BIN_CLIENT := "./bin/client"
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

generate:
	protoc --proto_path=api --go_out=internal/pb --go-grpc_out=internal/pb api/Service.proto

build-monitoring:
	go build -v -o $(BIN_MONITORING) -ldflags "$(LDFLAGS)" ./cmd/app

build-client:
	go build -v -o $(BIN_CLIENT) -ldflags "$(LDFLAGS)" ./cmd/client