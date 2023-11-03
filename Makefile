LOCAL_BIN:=$(CURDIR)/bin
PROJECT_VERSION=$(shell git describe)
REGISTRY_USER=sagata123
SERVICE_NAME=auth

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3

lint:
	GOBIN=$(LOCAL_BIN) golangci-lint run ./... --config .golangci.pipeline.yaml

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	make generate-user-api

GENERATED_OUTPUT_DIR:=$(CURDIR)/pkg/user_v1

generate-user-api:
	mkdir -p $(GENERATED_OUTPUT_DIR)
	protoc --proto_path api/user_v1 \
	--go_out=$(GENERATED_OUTPUT_DIR) --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=$(GENERATED_OUTPUT_DIR) --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/user_v1/user.proto

build:
	GOOS=linux GOARCH=amd64 go build -o $(LOCAL_BIN)/auth_$(PROJECT_VERSION) cmd/grpc_server/main.go

copy-to-server:
	scp $(LOCAL_BIN)/auth_$(PROJECT_VERSION) $(SERVER_USER)@$(SERVER_IP):/root

docker-build-and-push:
	docker buildx build --no-cache --platform linux/amd64 -t $(REGISTRY_DNS)/$(REGISTRY_USER)/$(SERVICE_NAME):$(PROJECT_VERSION) .
	docker login -u ${REGISTRY_USER} -p ${REGISTRY_PASS} $(REGISTRY_DNS)
	docker push $(REGISTRY_DNS)/$(REGISTRY_USER)/$(SERVICE_NAME):$(PROJECT_VERSION)