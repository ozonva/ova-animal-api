DIST_DIR = dist
GRPC_SERVICE_DIR = pkg/ova-animal-api

PGV_VERSION:="v0.6.1"
BUF_VERSION:="v0.51.0"
export GO111MODULE=on

.vendor-proto:
	mkdir -p vendor.protogen
	mkdir -p vendor.protogen/api
	cp api/animal.proto vendor.protogen/api/animal.proto
	@if [ ! -d vendor.protogen/google ]; then \
		git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
		mkdir -p  vendor.protogen/google/ &&\
		mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
		rm -rf vendor.protogen/googleapis;\
	fi
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.5.0
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.5.0
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go install github.com/envoyproxy/protoc-gen-validate@$(PGV_VERSION)


clean:
	rm -rf $(DIST_DIR)

fetch_dependencies:
	go mod download

generate: fetch_dependencies
	go get github.com/alvaroloes/enumer
	go install github.com/alvaroloes/enumer
	go install github.com/golang/mock/mockgen@v1.6.0
	go generate ./...

generate_proto: .vendor-proto
	mkdir -p $(GRPC_SERVICE_DIR)
	protoc -I vendor.protogen \
        --go_out=$(GRPC_SERVICE_DIR) --go_opt=paths=import \
        --go-grpc_out=$(GRPC_SERVICE_DIR) --go-grpc_opt=paths=import \
        --grpc-gateway_out=$(GRPC_SERVICE_DIR) --grpc-gateway_opt=logtostderr=true --grpc-gateway_opt=paths=import \
        --swagger_out=allow_merge=true,merge_file_name=api:$(GRPC_SERVICE_DIR) \
        --go-grpc_opt=paths=import api/animal.proto

test: fetch_dependencies generate_proto generate
	go test ./...

integration_test: fetch_dependencies generate_proto generate
	go test -tags=integration ./...

build: clean fetch_dependencies generate_proto generate
	mkdir $(DIST_DIR)
	go build -o $(DIST_DIR)/ ./...

run:
	go run ./...

all: test build
