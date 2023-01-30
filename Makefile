all: clean test

build:
	go mod tidy; go mod verify
	go build -o server/server server/main.go
	go build -o client/client client/main.go

test: build
	@echo "***** UNIT TESTS NOT YET PROVIDED *****"
	server/server &
	client/client &

clean:
	rm -f client/client server/server testtest *.log

.PHONY: all build test clean