DIST_DIR = dist

clean:
	rm -rf $(DIST_DIR)

fetch_dependencies:
	go mod download

test: fetch_dependencies
	go test ./...

build: clean fetch_dependencies
	mkdir $(DIST_DIR)
	go build -o $(DIST_DIR)/ ./...

run:
	go run ./...

all: test build
