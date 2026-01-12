.PHONY: build run clean

build:
	go build -o ./bin/gonogram
run:
	go run .
clean:
	rm -rf ./bin
