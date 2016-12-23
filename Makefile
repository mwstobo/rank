default: build

build:
	go build

test:
	go test ./...

clean:
	go clean

install:
	go install
