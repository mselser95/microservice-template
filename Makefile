build:
	go build -o my-microservice ./cmd/main.go

setup-pre-commit:
	brew install pre-commit
	pre-commit install
	go install github.com/golang/mock/mockgen@v1.6.0


generate-proto:
	protoc --go_out=./internal --go-grpc_out=./internal proto/service.proto

generate-api:
	openapi-generator-cli generate \
	  -i ./api/openapi.yaml \
	  -g go \
	  --skip-validate-spec \
	  --skip-go-mod \
	  --package-name internal/api \
	  -o ./internal/api/

generate-gql:
	gqlgen generate

mockgen:
	go generate ./...

test:
	go test -v ./... -cover

start:
	go run ./cmd/start.go

.PHONY: build generate-proto generate-gql start
