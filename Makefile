include env/.env
export

LOCAL_BIN:=$(CURDIR)/bin
PATH:=$(LOCAL_BIN):$(PATH)

APP_BIN = build/app

depend:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@latest

all: build

build: clean $(APP_BIN)
	./build/app
.PHONY: build

$(APP_BIN):
	go build -o $(APP_BIN) cmd/go_project_template/main.go

clean:
	rm -rf build
	rm -rf bin
.PHONY: clean

test:
	go test ./...
.PHONY: test

migrate-up:
	migrate -path migrations -database $(POSTGRES_DSN) up
.PHONY: migrate-up

migrate-down:
	migrate -path migrations -database $(POSTGRES_DSN) down
.PHONY: migrate-down

mock:
	mkdir ./internal/repository/repository_mocks
	mockgen -source ./internal/repository/repository.go -package repository_mocks > ./internal/repository/repository_mocks/mocks.go
.PHONY: mock

bin:
	GOBIN=$(LOCAL_BIN) go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	GOBIN=$(LOCAL_BIN) go install github.com/golang/mock/mockgen@latest
.PHONY: bin

docker.start:
	cd ./deployments && docker-compose up
.PHONY: docker.start

docker.stop:
	cd ./deployments && docker-compose stop && docker-compose rm -f
.PHONY: docker.stop

api:
	buf generate
.PHONY: api