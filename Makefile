# Target to simply run the go run command in the backend directory

.PHONY: run
run:
	go run .

build:
	go build

clean:
	go clean

.PHONY: test
test:
	go test -v ./...
