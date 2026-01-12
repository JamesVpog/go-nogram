.PHONY: build run clean

build:
	go build -o ./bin/gonogram
run:
	go run ./...
test:
	go test ./...
clean:
	rm -rf ./bin
