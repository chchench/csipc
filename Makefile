all: clean build

build:
	go mod tidy; go mod verify
	go build -o client client.go
	go build -o server-monitor server-monitor.go

clean:
	rm -f client server-monitor

.PHONY: all build clean
