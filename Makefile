DIST_DIR = dist

clean:
	rm -rf $(DIST_DIR)

fetch_dependencies:
	go mod download

generate: fetch_dependencies
	go get github.com/alvaroloes/enumer
	go install github.com/alvaroloes/enumer
	go install github.com/golang/mock/mockgen@v1.6.0
	go generate ./...

test: fetch_dependencies generate
	go test ./...

build: clean fetch_dependencies generate
	mkdir $(DIST_DIR)
	go build -o $(DIST_DIR)/ ./...

run:
	go run ./...

all: test build
