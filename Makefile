DIST_DIR = dist

clean:
	rm -rf $(DIST_DIR)

test:
	go test ./...

build: clean
	mkdir $(DIST_DIR)
	go build -o $(DIST_DIR)/ ./...

run:
	go run ./...

all: test build
